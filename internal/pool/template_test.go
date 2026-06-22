package pool

import "testing"

func TestJobToTemplateNil(t *testing.T) {
	var j *Job

	_, err := j.ToTemplate()
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestJobToTemplate(t *testing.T) {
	job := &Job{
		Height:   1,
		PrevHash: "0000000000000000000000000000000000000000000000000000000000000000",
		Version:  "20000000",
		NBits:    "1d00ffff",
		NTime:    "67d6b70d",
	}

	tpl, err := job.ToTemplate()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if tpl.Height != 1 {
		t.Fatalf("expected height 1 got %d", tpl.Height)
	}
}
