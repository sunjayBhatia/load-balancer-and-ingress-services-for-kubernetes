// Copyright 2021 VMware, Inc.
// SPDX-License-Identifier: Apache License 2.0
package models

// This file is auto-generated.

// WafLog waf log
// swagger:model WafLog
type WafLog struct {

	// Set to true if there are allowlist rules in the policy. Field introduced in 20.1.3. Allowed in Enterprise edition with any value, Enterprise with Cloud Services edition.
	AllowlistConfigured *bool `json:"allowlist_configured,omitempty"`

	// Log Entries generated by WAF allowlist rules. Field introduced in 20.1.3. Allowed in Enterprise edition with any value, Enterprise with Cloud Services edition.
	AllowlistLogs []*WafAllowlistLog `json:"allowlist_logs,omitempty"`

	// Set to true if allowlist rules were processed. Field introduced in 20.1.3. Allowed in Enterprise edition with any value, Enterprise with Cloud Services edition.
	AllowlistProcessed *bool `json:"allowlist_processed,omitempty"`

	// Log Entries generated by Application Specific Signature rules. Field introduced in 20.1.1. Allowed in Enterprise edition with any value, Essentials, Basic, Enterprise with Cloud Services edition.
	ApplicationRuleLogs []*WafRuleLog `json:"application_rule_logs,omitempty"`

	// Set to true if there are Application Specific Signature rules in the policy. Field introduced in 20.1.1. Allowed in Enterprise edition with any value, Essentials, Basic, Enterprise with Cloud Services edition.
	ApplicationRulesConfigured *bool `json:"application_rules_configured,omitempty"`

	// Set to true if Application Specific Signature rules were processed. Field introduced in 20.1.3. Allowed in Enterprise edition with any value, Enterprise with Cloud Services edition.
	ApplicationRulesProcessed *bool `json:"application_rules_processed,omitempty"`

	// Latency (in microseconds) in WAF Request Body Phase. Field introduced in 17.2.2. Allowed in Enterprise edition with any value, Essentials, Basic, Enterprise with Cloud Services edition.
	LatencyRequestBodyPhase *int64 `json:"latency_request_body_phase,omitempty"`

	// Latency (in microseconds) in WAF Request Header Phase. Field introduced in 17.2.2. Allowed in Enterprise edition with any value, Essentials, Basic, Enterprise with Cloud Services edition.
	LatencyRequestHeaderPhase *int64 `json:"latency_request_header_phase,omitempty"`

	// Latency (in microseconds) in WAF Response Body Phase. Field introduced in 17.2.2. Allowed in Enterprise edition with any value, Essentials, Basic, Enterprise with Cloud Services edition.
	LatencyResponseBodyPhase *int64 `json:"latency_response_body_phase,omitempty"`

	// Latency (in microseconds) in WAF Response Header Phase. Field introduced in 17.2.2. Allowed in Enterprise edition with any value, Essentials, Basic, Enterprise with Cloud Services edition.
	LatencyResponseHeaderPhase *int64 `json:"latency_response_header_phase,omitempty"`

	// The total memory (in bytes) consumed by WAF to process this request. Field introduced in 22.1.1. Unit is BYTES. Allowed in Enterprise edition with any value, Enterprise with Cloud Services edition.
	MemoryAllocated *int64 `json:"memory_allocated,omitempty"`

	// Omitted Application rule log stats. Field introduced in 22.1.1. Allowed in Enterprise edition with any value, Enterprise with Cloud Services edition.
	OmittedAppRuleStats *OmittedWafLogStats `json:"omitted_app_rule_stats,omitempty"`

	// Omitted WAF rule log stats. Field introduced in 22.1.1. Allowed in Enterprise edition with any value, Enterprise with Cloud Services edition.
	OmittedSignatureStats *OmittedWafLogStats `json:"omitted_signature_stats,omitempty"`

	// Set to true if there are Positive Security Model rules in the policy. Field introduced in 18.2.3. Allowed in Enterprise edition with any value, Essentials, Basic, Enterprise with Cloud Services edition.
	PsmConfigured *bool `json:"psm_configured,omitempty"`

	// Log Entries generated by WAF Positive Security Model. Field introduced in 18.2.3. Allowed in Enterprise edition with any value, Essentials, Basic, Enterprise with Cloud Services edition.
	PsmLogs []*WafPSMLog `json:"psm_logs,omitempty"`

	// Set to true if Positive Security Model rules were processed. Field introduced in 20.1.3. Allowed in Enterprise edition with any value, Enterprise with Cloud Services edition.
	PsmProcessed *bool `json:"psm_processed,omitempty"`

	//  Field introduced in 17.2.1. Allowed in Enterprise edition with any value, Essentials, Basic, Enterprise with Cloud Services edition.
	RuleLogs []*WafRuleLog `json:"rule_logs,omitempty"`

	// Set to true if there are ModSecurity rules in the policy. Field introduced in 18.2.3. Allowed in Enterprise edition with any value, Essentials, Basic, Enterprise with Cloud Services edition.
	RulesConfigured *bool `json:"rules_configured,omitempty"`

	// Set to true if ModSecurity rules were processed. Field introduced in 20.1.3. Allowed in Enterprise edition with any value, Enterprise with Cloud Services edition.
	RulesProcessed *bool `json:"rules_processed,omitempty"`

	// Denotes whether WAF is running in detection mode or enforcement mode, whether any rules matched the transaction, and whether transaction is dropped by the WAF module. Enum options - NO_WAF, FLAGGED, PASSED, REJECTED, ALLOWLISTED, BYPASSED. Field introduced in 17.2.2. Allowed in Enterprise edition with any value, Essentials, Basic, Enterprise with Cloud Services edition.
	Status *string `json:"status,omitempty"`
}
