// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// DeviceManagementScheduledReportRecurrence undocumented
type DeviceManagementScheduledReportRecurrence int

const (
	// DeviceManagementScheduledReportRecurrenceVNone undocumented
	DeviceManagementScheduledReportRecurrenceVNone DeviceManagementScheduledReportRecurrence = 0
	// DeviceManagementScheduledReportRecurrenceVDaily undocumented
	DeviceManagementScheduledReportRecurrenceVDaily DeviceManagementScheduledReportRecurrence = 1
	// DeviceManagementScheduledReportRecurrenceVWeekly undocumented
	DeviceManagementScheduledReportRecurrenceVWeekly DeviceManagementScheduledReportRecurrence = 2
	// DeviceManagementScheduledReportRecurrenceVMonthly undocumented
	DeviceManagementScheduledReportRecurrenceVMonthly DeviceManagementScheduledReportRecurrence = 3
)

// DeviceManagementScheduledReportRecurrencePNone returns a pointer to DeviceManagementScheduledReportRecurrenceVNone
func DeviceManagementScheduledReportRecurrencePNone() *DeviceManagementScheduledReportRecurrence {
	v := DeviceManagementScheduledReportRecurrenceVNone
	return &v
}

// DeviceManagementScheduledReportRecurrencePDaily returns a pointer to DeviceManagementScheduledReportRecurrenceVDaily
func DeviceManagementScheduledReportRecurrencePDaily() *DeviceManagementScheduledReportRecurrence {
	v := DeviceManagementScheduledReportRecurrenceVDaily
	return &v
}

// DeviceManagementScheduledReportRecurrencePWeekly returns a pointer to DeviceManagementScheduledReportRecurrenceVWeekly
func DeviceManagementScheduledReportRecurrencePWeekly() *DeviceManagementScheduledReportRecurrence {
	v := DeviceManagementScheduledReportRecurrenceVWeekly
	return &v
}

// DeviceManagementScheduledReportRecurrencePMonthly returns a pointer to DeviceManagementScheduledReportRecurrenceVMonthly
func DeviceManagementScheduledReportRecurrencePMonthly() *DeviceManagementScheduledReportRecurrence {
	v := DeviceManagementScheduledReportRecurrenceVMonthly
	return &v
}