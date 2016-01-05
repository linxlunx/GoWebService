package helpers

func StatusList() T {
	resp := map[int]interface{}{
		100: "parameters is not complete",
		404: "data not found",
	}

	return T(resp)
}
