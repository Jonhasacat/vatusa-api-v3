package notifications

type NotificationTemplate string

// Notifications sent to controllers/users
const (
	ControllerDivisionJoin           NotificationTemplate = "controller_division_join"
	ControllerTransferRequested      NotificationTemplate = "controller_transfer_requested"
	ControllerTransferAccepted       NotificationTemplate = "controller_transfer_accepted"
	ControllerVisitingRequested      NotificationTemplate = "controller_visiting_requested"
	ControllerVisitorAccepted        NotificationTemplate = "controller_visitor_accepted"
	ControllerHomeRemoved            NotificationTemplate = "controller_home_removed"
	ControllerVisitorRemoved         NotificationTemplate = "controller_visitor_removed"
	ControllerAcademyCourseAssigned  NotificationTemplate = "controller_academy_course_assigned"
	ControllerAcademyCourseCompleted NotificationTemplate = "controller_academy_course_completed"
	ControllerAcademyExamAssigned    NotificationTemplate = "controller_academy_exam_assigned"
	ControllerAcademyExamPassed      NotificationTemplate = "controller_academy_exam_passed"
	ControllerAcademyExamFailed      NotificationTemplate = "controller_academy_exam_failed"
)

// Notifications sent to staff
const (
	FacilityTransferRequested     NotificationTemplate = "facility_transfer_requested"
	FacilityTransferAccepted      NotificationTemplate = "facility_transfer_accepted"
	FacilityVisitingRequested     NotificationTemplate = "facility_visiting_requested"
	FacilityRosterRequestsOverdue NotificationTemplate = "facility_roster_requests_overdue"
	FacilityHomeRemoved           NotificationTemplate = "facility_home_removed"
	FacilityVisitorRemoved        NotificationTemplate = "facility_visitor_removed"
)

// Notifications sent to division staff

const ()
