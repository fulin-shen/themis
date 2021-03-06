package policy

import (
	"fmt"
	"net"
	"testing"

	"github.com/infobloxopen/themis/contrib/coredns/policy/testutil"
	"github.com/infobloxopen/themis/pdp"
	"github.com/miekg/dns"
)

func TestPatchDebugMsg(t *testing.T) {
	p := newPolicyPlugin()
	p.conf.debugSuffix = "debug.local."

	m := testutil.MakeTestDNSMsg("example.com.debug.local", dns.TypeTXT, dns.ClassCHAOS)
	dbgMsgr := p.patchDebugMsg(m)
	if dbgMsgr == nil {
		t.Error("expected patched message")
	}

	testutil.AssertDNSMessage(t, "patchDebugMsg", 0, m, 0,
		";; opcode: QUERY, status: NOERROR, id: 0\n"+
			";; flags:; QUERY: 1, ANSWER: 0, AUTHORITY: 0, ADDITIONAL: 0\n\n"+
			";; QUESTION SECTION:\n"+
			";example.com.\tIN\t A\n",
	)

	m = new(dns.Msg)
	dbgMsgr = p.patchDebugMsg(m)
	if dbgMsgr != nil {
		t.Error("expected NOT patched message")
	}

	testutil.AssertDNSMessage(t, "patchDebugMsg(nil)", 0, m, 0,
		";; opcode: QUERY, status: NOERROR, id: 0\n"+
			";; flags:; QUERY: 0, ANSWER: 0, AUTHORITY: 0, ADDITIONAL: 0\n",
	)

	m = testutil.MakeTestDNSMsg("example.com", dns.TypeTXT, dns.ClassCHAOS)
	dbgMsgr = p.patchDebugMsg(m)
	if dbgMsgr != nil {
		t.Error("expected NOT patched message")
	}

	testutil.AssertDNSMessage(t, "patchDebugMsg(nil)", 0, m, 0,
		";; opcode: QUERY, status: NOERROR, id: 0\n"+
			";; flags:; QUERY: 1, ANSWER: 0, AUTHORITY: 0, ADDITIONAL: 0\n\n"+
			";; QUESTION SECTION:\n"+
			";example.com.\tCH\t TXT\n",
	)
}

func TestNewDbgMessenger(t *testing.T) {
	dm := newDbgMessenger(testOrgName, "id")
	if dm.orgName != testOrgName {
		t.Errorf("Unexpected original name: %s", dm.orgName)
	}
	if dm.String() != "Ident: id" {
		t.Errorf("Unexpected init message: %s", dm.String())
	}

	dm = newDbgMessenger(testOrgName, "")
	if dm.orgName != testOrgName {
		t.Errorf("Unexpected original name: %s", dm.orgName)
	}
	if dm.String() != "" {
		t.Errorf("Unexpected init message: %s", dm.String())
	}
}

func TestTxtMsgs(t *testing.T) {
	dm := newTestDbgMessenger()
	dm.WriteString("extremely long string, 25extremely long string, 50")
	dm.WriteString("extremely long string, 75extremely long string,100")
	dm.WriteString("extremely long string,125extremely long string,150")
	dm.WriteString("extremely long string,175extremely long string,200")
	dm.WriteString("extremely long string,225extremely long string,250")
	dm.WriteString("extremely long string,275extremely long string,300")
	dm.msgBounds = []int{7, 296, 300}

	txts := dm.txtMsgs()
	if len(txts) != 4 {
		t.Errorf("Unexpected number of TXT messages: %q", txts)
		return
	}
	for i, exp := range []string{
		"extreme",
		"ly long string, 25extremely long string, 50" +
			"extremely long string, 75extremely long string,100" +
			"extremely long string,125extremely long string,150" +
			"extremely long string,175extremely long string,200" +
			"extremely long string,225extremely long string,250" +
			"extremely lo",
		"ng string,275extremely long string",
		",300",
	} {
		if txts[i] != exp {
			t.Errorf("Unexpected TXT #%d: expected %q, actual %q", i, exp, txts[i])
		}
	}
}

func TestDebugNameVal(t *testing.T) {
	n, v := debugNameVal(pdp.MakeStringAssignment("attr1", "value1"))
	if n != "attr1" || v != "value1" {
		t.Errorf("Unexpected result: name=%s, val=%s", n, v)
	}

	n, v = debugNameVal(pdp.MakeIntegerAssignment("attr2", 12345))
	if n != "attr2" || v != "12345" {
		t.Errorf("Unexpected result: name=%s, val=%s", n, v)
	}

	n, v = debugNameVal(pdp.MakeAddressAssignment("attr3", net.ParseIP("2001:db8::1")))
	if n != "attr3" || v != "2001:db8::1" {
		t.Errorf("Unexpected result: name=%s, val=%s", n, v)
	}
}

func TestAppendAttrs(t *testing.T) {
	dm := newTestDbgMessenger()
	dm.appendAttrs("attributes", []pdp.AttributeAssignment{
		pdp.MakeStringAssignment("attr1", "value1"),
		pdp.MakeIntegerAssignment("attr2", 12345),
		pdp.MakeAddressAssignment("attr3", net.ParseIP("2001:db8::1")),
	})
	if dm.String() != "attributes: [attr1: value1, attr2: 12345, attr3: 2001:db8::1]" {
		t.Errorf("Unexpected result: %s", dm.String())
	}
}

func TestAppendPassthrough(t *testing.T) {
	dm := newTestDbgMessenger()
	dm.appendPassthrough()
	if dm.String() != "Passthrough: yes" {
		t.Errorf("Unexpected result: %s", dm.String())
	}
}

func TestAppendDebugID(t *testing.T) {
	dm := newTestDbgMessenger()
	dm.appendDebugID("DeBuGiD")
	if dm.String() != "Ident: DeBuGiD" {
		t.Errorf("Unexpected result: %s", dm.String())
	}
}

func TestAppendResolution(t *testing.T) {
	dm := newTestDbgMessenger()
	dm.appendResolution(-1)
	if dm.String() != "Domain resolution: skipped" {
		t.Errorf("Unexpected result: %s", dm.String())
	}

	dm = newTestDbgMessenger()
	dm.appendResolution(dns.RcodeSuccess)
	if dm.String() != "Domain resolution: resolved" {
		t.Errorf("Unexpected result: %s", dm.String())
	}

	dm = newTestDbgMessenger()
	dm.appendResolution(dns.RcodeServerFailure)
	if dm.String() != "Domain resolution: failed" {
		t.Errorf("Unexpected result: %s", dm.String())
	}

	dm = newTestDbgMessenger()
	dm.appendResolution(dns.RcodeNameError)
	if dm.String() != "Domain resolution: not resolved" {
		t.Errorf("Unexpected result: %s", dm.String())
	}
}

func TestAppendResponse(t *testing.T) {
	dm := newTestDbgMessenger()
	r := pdp.Response{
		Effect: pdp.EffectPermit,
		Obligations: []pdp.AttributeAssignment{
			pdp.MakeStringAssignment("attr1", "value1"),
			pdp.MakeIntegerAssignment("attr2", 12345),
			pdp.MakeAddressAssignment("attr3", net.ParseIP("2001:db8::1")),
		},
	}
	dm.appendResponse(&r)
	if dm.String() != "PDP response {Effect: Permit, Obligations: [attr1: value1, attr2: 12345, attr3: 2001:db8::1]}" {
		t.Errorf("Unexpected result: %s", dm.String())
	}
}

func TestSetDebugQueryPassthroughAnswer(t *testing.T) {
	dm := newDbgMessenger("example.passthrough.local.debug.local.", "")
	m := testutil.MakeTestDNSMsg("example.passthrough.local", dns.TypeA, dns.ClassINET)

	dm.setDebugQueryPassthroughAnswer(m)
	testutil.AssertDNSMessage(t, "setDebugQueryPassthroughAnswer", 0, m, 0,
		";; opcode: QUERY, status: NOERROR, id: 0\n"+
			";; flags:; QUERY: 1, ANSWER: 1, AUTHORITY: 0, ADDITIONAL: 0\n\n"+
			";; QUESTION SECTION:\n"+
			";example.passthrough.local.\tIN\t A\n\n"+
			";; ANSWER SECTION:\n"+
			"example.passthrough.local.debug.local.\t0\tCH\tTXT\t\"Passthrough: yes\"\n",
	)
}

func TestSetDebugQueryAnswer(t *testing.T) {
	t.Run("OnDomainResponse(RcodeSuccess)", func(t *testing.T) {
		p := newPolicyPlugin()
		p.conf.debugSuffix = "debug.local."
		p.conf.debugID = "<DEBUG>"

		mpc := testutil.NewMockPdpClient(t)
		p.pdp = mpc

		m := testutil.MakeTestDNSMsg("example.com.debug.local.", dns.TypeTXT, dns.ClassCHAOS)
		w := testutil.NewTestAddressedNonwriter("192.0.2.1")
		dm := p.patchDebugMsg(m)

		cfg := newConfig()
		cfg.attrs.parseAttrList(attrListTypeVal1, "domain_name")

		ah := newAttrHolder(nil, cfg.attrs)
		ah.addDnsQuery(w, m, cfg.options)
		mpc.Out = []pdp.AttributeAssignment{
			pdp.MakeStringAssignment(attrNameRedirectTo, "192.0.2.54"),
			pdp.MakeIntegerAssignment("policy_action", 4),
		}
		mpc.Effect = pdp.EffectPermit

		p.validate(nil, ah, attrListTypeVal1, dm)

		dm.setDebugQueryAnswer(m, dns.RcodeSuccess)
		testutil.AssertDNSMessage(t, "setDebugQueryAnswer", 0, m, 0,
			";; opcode: QUERY, status: NOERROR, id: 0\n"+
				";; flags:; QUERY: 1, ANSWER: 1, AUTHORITY: 0, ADDITIONAL: 0\n\n"+
				";; QUESTION SECTION:\n"+
				";example.com.\tIN\t A\n\n"+
				";; ANSWER SECTION:\n"+
				"example.com.debug.local.\t0\tCH\tTXT\t\"Ident: <DEBUG>\" "+
				"\"PDP response {Effect: Permit, Obligations: [redirect_to: 192.0.2.54, policy_action: redirect]}\" "+
				"\"Domain resolution: resolved\"\n",
		)
	})

	t.Run("OnDomainResponse(RcodeNameError)", func(t *testing.T) {
		p := newPolicyPlugin()
		p.conf.debugSuffix = "debug.local."
		p.conf.debugID = "<DEBUG>"

		mpc := testutil.NewMockPdpClient(t)
		p.pdp = mpc

		m := testutil.MakeTestDNSMsg("example.com.debug.local.", dns.TypeTXT, dns.ClassCHAOS)
		w := testutil.NewTestAddressedNonwriter("192.0.2.1")
		dm := p.patchDebugMsg(m)

		cfg := newConfig()
		cfg.attrs.parseAttrList(attrListTypeVal1, "domain_name")

		ah := newAttrHolder(nil, cfg.attrs)
		ah.addDnsQuery(w, m, cfg.options)
		mpc.Out = []pdp.AttributeAssignment{
			pdp.MakeIntegerAssignment("policy_action", 2),
		}
		mpc.Effect = pdp.EffectPermit

		p.validate(nil, ah, attrListTypeVal1, dm)

		dm.setDebugQueryAnswer(m, dns.RcodeNameError)
		testutil.AssertDNSMessage(t, "setDebugQueryAnswer", 0, m, 0,
			";; opcode: QUERY, status: NOERROR, id: 0\n"+
				";; flags:; QUERY: 1, ANSWER: 1, AUTHORITY: 0, ADDITIONAL: 0\n\n"+
				";; QUESTION SECTION:\n"+
				";example.com.\tIN\t A\n\n"+
				";; ANSWER SECTION:\n"+
				"example.com.debug.local.\t0\tCH\tTXT\t\"Ident: <DEBUG>\" "+
				"\"PDP response {Effect: Permit, Obligations: [policy_action: allow]}\" "+
				"\"Domain resolution: not resolved\"\n",
		)
	})

	t.Run("OnDomainResponse(RcodeServerFailure)", func(t *testing.T) {
		p := newPolicyPlugin()
		p.conf.debugSuffix = "debug.local."
		p.conf.debugID = "<DEBUG>"

		mpc := testutil.NewMockPdpClient(t)
		p.pdp = mpc

		m := testutil.MakeTestDNSMsg("example.com.debug.local.", dns.TypeTXT, dns.ClassCHAOS)
		w := testutil.NewTestAddressedNonwriter("192.0.2.1")
		dm := p.patchDebugMsg(m)

		cfg := newConfig()
		cfg.attrs.parseAttrList(attrListTypeVal1, "domain_name")

		ah := newAttrHolder(nil, cfg.attrs)
		ah.addDnsQuery(w, m, cfg.options)
		mpc.Out = []pdp.AttributeAssignment{
			pdp.MakeIntegerAssignment("policy_action", 2),
		}
		mpc.Effect = pdp.EffectPermit

		p.validate(nil, ah, attrListTypeVal1, dm)

		dm.setDebugQueryAnswer(m, dns.RcodeServerFailure)
		testutil.AssertDNSMessage(t, "setDebugQueryAnswer", 0, m, 0,
			";; opcode: QUERY, status: NOERROR, id: 0\n"+
				";; flags:; QUERY: 1, ANSWER: 1, AUTHORITY: 0, ADDITIONAL: 0\n\n"+
				";; QUESTION SECTION:\n"+
				";example.com.\tIN\t A\n\n"+
				";; ANSWER SECTION:\n"+
				"example.com.debug.local.\t0\tCH\tTXT\t\"Ident: <DEBUG>\" "+
				"\"PDP response {Effect: Permit, Obligations: [policy_action: allow]}\" "+
				"\"Domain resolution: failed\"\n",
		)
	})

	t.Run("OnDomainResponse(-1)", func(t *testing.T) {
		p := newPolicyPlugin()
		p.conf.debugSuffix = "debug.local."
		p.conf.debugID = "<DEBUG>"

		mpc := testutil.NewMockPdpClient(t)
		p.pdp = mpc

		m := testutil.MakeTestDNSMsg("example.com.debug.local.", dns.TypeTXT, dns.ClassCHAOS)
		w := testutil.NewTestAddressedNonwriter("192.0.2.1")
		dm := p.patchDebugMsg(m)

		cfg := newConfig()
		cfg.attrs.parseAttrList(attrListTypeVal1, "domain_name")

		ah := newAttrHolder(nil, cfg.attrs)
		ah.addDnsQuery(w, m, cfg.options)
		mpc.Out = []pdp.AttributeAssignment{
			pdp.MakeIntegerAssignment("policy_action", 2),
		}
		mpc.Effect = pdp.EffectPermit

		p.validate(nil, ah, attrListTypeVal1, dm)

		dm.setDebugQueryAnswer(m, -1)
		testutil.AssertDNSMessage(t, "setDebugQueryAnswer", 0, m, 0,
			";; opcode: QUERY, status: NOERROR, id: 0\n"+
				";; flags:; QUERY: 1, ANSWER: 1, AUTHORITY: 0, ADDITIONAL: 0\n\n"+
				";; QUESTION SECTION:\n"+
				";example.com.\tIN\t A\n\n"+
				";; ANSWER SECTION:\n"+
				"example.com.debug.local.\t0\tCH\tTXT\t\"Ident: <DEBUG>\" "+
				"\"PDP response {Effect: Permit, Obligations: [policy_action: allow]}\" "+
				"\"Domain resolution: skipped\"\n",
		)
	})

	t.Run("OnIPResponse", func(t *testing.T) {
		p := newPolicyPlugin()
		p.conf.debugSuffix = "debug.local."
		p.conf.debugID = "<DEBUG>"

		mpc := testutil.NewMockPdpClient(t)
		p.pdp = mpc

		m := testutil.MakeTestDNSMsg("example.com.debug.local.", dns.TypeTXT, dns.ClassCHAOS)
		w := testutil.NewTestAddressedNonwriter("192.0.2.1")
		dm := p.patchDebugMsg(m)

		cfg := newConfig()
		cfg.attrs.parseAttrList(attrListTypeVal1, "domain_name")
		cfg.attrs.parseAttrList(attrListTypeVal2, "address")

		ah := newAttrHolder(nil, cfg.attrs)
		ah.addDnsQuery(w, m, cfg.options)
		mpc.Out = []pdp.AttributeAssignment{
			pdp.MakeIntegerAssignment("policy_action", 2),
		}
		mpc.Effect = pdp.EffectPermit

		p.validate(nil, ah, attrListTypeVal1, dm)

		ah.addAddressAttr(net.ParseIP("2001:db8::1"))
		mpc.Out = []pdp.AttributeAssignment{
			pdp.MakeIntegerAssignment("policy_action", 4),
			pdp.MakeStringAssignment(attrNameRedirectTo, "192.0.2.54"),
		}

		p.validate(nil, ah, attrListTypeVal2, dm)

		dm.setDebugQueryAnswer(m, dns.RcodeSuccess)
		testutil.AssertDNSMessage(t, "setDebugQueryAnswer", 0, m, 0,
			";; opcode: QUERY, status: NOERROR, id: 0\n"+
				";; flags:; QUERY: 1, ANSWER: 1, AUTHORITY: 0, ADDITIONAL: 0\n\n"+
				";; QUESTION SECTION:\n"+
				";example.com.\tIN\t A\n\n"+
				";; ANSWER SECTION:\n"+
				"example.com.debug.local.\t0\tCH\tTXT\t\"Ident: <DEBUG>\" "+
				"\"PDP response {Effect: Permit, Obligations: [policy_action: allow]}\" "+
				"\"PDP response {Effect: Permit, Obligations: [policy_action: redirect, redirect_to: 192.0.2.54]}\" "+
				"\"Domain resolution: resolved\"\n",
		)
	})

	t.Run("DefaultDecisionOnError", func(t *testing.T) {
		p := newPolicyPlugin()
		p.conf.debugSuffix = "debug.local."
		p.conf.debugID = "<DEBUG>"

		mpc := testutil.NewMockPdpClient(t)
		p.pdp = mpc

		m := testutil.MakeTestDNSMsg("example.com.debug.local.", dns.TypeTXT, dns.ClassCHAOS)
		w := testutil.NewTestAddressedNonwriter("192.0.2.1")
		dm := p.patchDebugMsg(m)

		cfg := newConfig()
		cfg.attrs.parseAttrList(attrListTypeVal1, "domain_name")
		cfg.attrs.parseAttrList(attrListTypeDefDecision, "policy_action=2")

		ah := newAttrHolder(nil, cfg.attrs)
		ah.addDnsQuery(w, m, cfg.options)
		mpc.Err = fmt.Errorf("test PDP error")
		mpc.Effect = pdp.EffectPermit

		p.validate(nil, ah, attrListTypeVal1, dm)

		dm.setDebugQueryAnswer(m, dns.RcodeSuccess)
		testutil.AssertDNSMessage(t, "setDebugQueryAnswer", 0, m, 0,
			";; opcode: QUERY, status: NOERROR, id: 0\n"+
				";; flags:; QUERY: 1, ANSWER: 1, AUTHORITY: 0, ADDITIONAL: 0\n\n"+
				";; QUESTION SECTION:\n"+
				";example.com.\tIN\t A\n\n"+
				";; ANSWER SECTION:\n"+
				"example.com.debug.local.\t0\tCH\tTXT\t\"Ident: <DEBUG>\" "+
				"\"Default decision: [policy_action: allow]\" "+
				"\"Domain resolution: resolved\"\n",
		)
	})

	t.Run("DefaultDecisionOnEffectIndeterminate", func(t *testing.T) {
		p := newPolicyPlugin()
		p.conf.debugSuffix = "debug.local."
		p.conf.debugID = "<DEBUG>"

		mpc := testutil.NewMockPdpClient(t)
		p.pdp = mpc

		m := testutil.MakeTestDNSMsg("example.com.debug.local.", dns.TypeTXT, dns.ClassCHAOS)
		w := testutil.NewTestAddressedNonwriter("192.0.2.1")
		dm := p.patchDebugMsg(m)

		cfg := newConfig()
		cfg.attrs.parseAttrList(attrListTypeVal1, "domain_name")
		cfg.attrs.parseAttrList(attrListTypeDefDecision, "policy_action=2")

		ah := newAttrHolder(nil, cfg.attrs)
		ah.addDnsQuery(w, m, cfg.options)
		mpc.Status = fmt.Errorf("test PDP error")
		mpc.Effect = pdp.EffectIndeterminate

		p.validate(nil, ah, attrListTypeVal1, dm)

		dm.setDebugQueryAnswer(m, dns.RcodeSuccess)
		testutil.AssertDNSMessage(t, "setDebugQueryAnswer", 0, m, 0,
			";; opcode: QUERY, status: NOERROR, id: 0\n"+
				";; flags:; QUERY: 1, ANSWER: 1, AUTHORITY: 0, ADDITIONAL: 0\n\n"+
				";; QUESTION SECTION:\n"+
				";example.com.\tIN\t A\n\n"+
				";; ANSWER SECTION:\n"+
				"example.com.debug.local.\t0\tCH\tTXT\t\"Ident: <DEBUG>\" "+
				"\"PDP response {Effect: Indeterminate, Obligations: []}\" "+
				"\"Default decision: [policy_action: allow]\" "+
				"\"Domain resolution: resolved\"\n",
		)
	})
}

const testOrgName = "test.com.debug."

func newTestDbgMessenger() *dbgMessenger {
	return &dbgMessenger{orgName: testOrgName}
}
