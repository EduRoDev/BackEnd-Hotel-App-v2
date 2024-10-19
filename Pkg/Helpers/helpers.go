package helpers

func Error(err error, message string) map[string]interface{} {
	return map[string]interface{}{"error": err, "message": message}
}

func ErrorWithStatus(err string, message string, status string) map[string]interface{} {
	return map[string]interface{}{"error": err, "message": message, "status": status}
}

func Success(message string) map[string]interface{} {
	return map[string]interface{}{"success": true, "message": message}
}


