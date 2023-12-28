package science

import (
	"encoding/json"
	"sort"
	"strings"
)

type agencyT struct {
	Agencies []reportT
}

type reportT struct {
	SubmittedResearchReports []saveDataT
}

type saveDataT struct {
	ExperimentID       string
	ResearchLocationID string
	ResearchReportType string
	FinalScienceValue  float32
}

type SortedT struct {
	Body map[string]ReportT
}

type ReportT struct {
	Data []DataT
}

type DataT struct {
	State       string
	Biome       string
	Description string
	Type        string
	Value       int
}

// **************************************************** Setters ********************************************************

func (s *SortedT) SetData(raw []byte) error {
	var agency agencyT
	if err := json.Unmarshal(raw, &agency); err != nil {
		return err
	}
	s.sortOnBody(agency.Agencies[0].SubmittedResearchReports)
	return nil
}

func (s *SortedT) sortOnBody(in []saveDataT) {
	for _, v := range in {
		location := strings.Split(v.ResearchLocationID, "_")
		biome := "n/a"
		if len(location) == 3 {
			biome, _ = strings.CutPrefix(location[2], location[0])
		}
		dType, _ := strings.CutSuffix(v.ResearchReportType, "Type")
		newData := DataT{
			State:       location[1],
			Biome:       biome,
			Description: v.ExperimentID,
			Type:        dType,
			Value:       int(v.FinalScienceValue),
		}
		if report, found := s.Body[location[0]]; found {
			report.Data = append(report.Data, newData)
			delete(s.Body, location[0])
			s.Body[location[0]] = report
		} else {
			var r ReportT
			r.Data = append(r.Data, newData)
			s.Body[location[0]] = r
		}
	}
	s.sortOnBiome()
}

func (s *SortedT) sortOnBiome() {
	for k, v := range s.Body {
		data := v.Data
		sort.SliceStable(data, func(i, j int) bool {
			return data[i].Biome < data[j].Biome
		})
		delete(s.Body, k)
		var r ReportT
		r.Data = data
		s.Body[k] = r
	}
}

// **************************************************** Getters ********************************************************
