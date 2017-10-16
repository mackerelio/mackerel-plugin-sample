package mpsample

import "testing"

func TestFetchMetrics(t *testing.T) {
	sample := &SamplePlugin{}

	ret, err := sample.FetchMetrics()

	if err != nil {
		t.Errorf("FetchMetrics returns error")
	}

	if !(1 <= ret["d6"] && ret["d6"] <= 6) {
		t.Errorf("FetchMetrics doesn't return a value between 1 and 6 for d6 metrics")
	}

	if !(1 <= ret["d20"] && ret["d20"] <= 20) {
		t.Errorf("FetchMetrics doesn't return a value between 1 and 20 for d20 metrics")
	}
}
