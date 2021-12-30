package speedtrackertypes

func CalculateSessionSpeedStatistics(session *Session) {
	for i := 0; i < len(session.Swings); i++ {
		session.TotalSpeed = session.TotalSpeed + session.Swings[i].Speed
	}

	session.AverageSpeed = session.TotalSpeed / len(session.Swings)
}