// AUTOGENERATED - DO NOT EDIT
// Generated by go-windows.

// Package taskscheduler implements the Windows.Win32.TaskScheduler namespace.
package taskscheduler

type DAILY struct {
	DaysInterval int
}

type WEEKLY struct {
	WeeksInterval    int
	RGFDAYSOFTHEWEEK int
}

type MONTHLYDATE struct {
	RGFDAYS   int
	RGFMONTHS int
}

type MONTHLYDOW struct {
	WWHICHWEEK       int
	RGFDAYSOFTHEWEEK int
	RGFMONTHS        int
}

type TRIGGER_TYPE_UNION struct {
	Daily       int
	Weekly      int
	MonthlyDate int
	MonthlyDOW  int
}

type TASK_TRIGGER struct {
	CBTRIGGERSIZE          int
	Reserved1              int
	WBEGINYEAR             int
	WBEGINMONTH            int
	WBEGINDAY              int
	WENDYEAR               int
	WENDMONTH              int
	WENDDAY                int
	WSTARTHOUR             int
	WSTARTMINUTE           int
	MinutesDuration        int
	MinutesInterval        int
	RGFLAGS                int
	TriggerType            int
	Type                   int
	Reserved2              int
	WRANDOMMINUTESINTERVAL int
}

type TaskScheduler struct {
}

type TaskHandlerPS struct {
}

type TaskHandlerStatusPS struct {
}