package speedtrackertypes

func CalculateSessionSpeedStatistics(session *Session) {
	session.MinimumSpeed = session.Swings[0].Speed
	session.MaximumSpeed = session.Swings[0].Speed

	for i := 0; i < len(session.Swings); i++ {
		session.TotalSpeed = session.TotalSpeed + session.Swings[i].Speed

		if session.Swings[i].Speed > session.MaximumSpeed {
			session.MaximumSpeed = session.Swings[i].Speed
		}

		if session.Swings[i].Speed < session.MinimumSpeed {
			session.MinimumSpeed = session.Swings[i].Speed
		}
	}

	session.AverageSpeed = session.TotalSpeed / len(session.Swings)
}
