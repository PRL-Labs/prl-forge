package stratumv2

import "log"

func HandleAuthorize(session *Session, msg map[string]any) error {

	log.Printf("🔐 mining.authorize")

	resp := map[string]any{
		"id":     msg["id"],
		"result": true,
		"error":  nil,
	}

	if err := session.Send(resp); err != nil {
		return err
	}

	diff := map[string]any{
		"id":     nil,
		"method": "mining.set_difficulty",
		"params": []any{1},
	}

	if err := session.Send(diff); err != nil {
		return err
	}

	log.Println("✅ Difficulty sent")

	return nil
}