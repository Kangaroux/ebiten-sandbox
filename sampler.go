package game

type Sampler struct {
	index      int
	maxSamples int
	samples    []float64
}

func NewSampler(maxSamples int) *Sampler {
	return &Sampler{
		maxSamples: maxSamples,
		samples:    make([]float64, maxSamples),
	}
}

func (s *Sampler) Add(val float64) {
	s.samples[s.index] = val
	s.index = (s.index + 1) % s.maxSamples
}

func (s *Sampler) Average() (avg float64) {
	for _, v := range s.samples {
		avg += v
	}

	return avg / float64(s.maxSamples)
}
