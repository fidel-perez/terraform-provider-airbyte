// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// ConnectionScheduleResponse - schedule for when the the connection should run, per the schedule type
type ConnectionScheduleResponse struct {
	BasicTiming    *string                   `json:"basicTiming,omitempty"`
	CronExpression *string                   `json:"cronExpression,omitempty"`
	ScheduleType   ScheduleTypeWithBasicEnum `json:"scheduleType"`
}

func (o *ConnectionScheduleResponse) GetBasicTiming() *string {
	if o == nil {
		return nil
	}
	return o.BasicTiming
}

func (o *ConnectionScheduleResponse) GetCronExpression() *string {
	if o == nil {
		return nil
	}
	return o.CronExpression
}

func (o *ConnectionScheduleResponse) GetScheduleType() ScheduleTypeWithBasicEnum {
	if o == nil {
		return ScheduleTypeWithBasicEnum("")
	}
	return o.ScheduleType
}
