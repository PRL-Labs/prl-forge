package block

// EncodeScriptNum encodes an integer using Bitcoin ScriptNum encoding (BIP34).
func EncodeScriptNum(v int64) []byte {
	if v == 0 {
		return nil
	}

	neg := v < 0
	if neg {
		v = -v
	}

	var out []byte

	for v > 0 {
		out = append(out, byte(v))
		v >>= 8
	}

	if out[len(out)-1]&0x80 != 0 {
		if neg {
			out = append(out, 0x80)
		} else {
			out = append(out, 0x00)
		}
	} else if neg {
		out[len(out)-1] |= 0x80
	}

	return out
}