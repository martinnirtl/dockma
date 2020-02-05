package dockercompose

import "sort"

func getServicesFromStringMap(serviceMap map[string]interface{}) []string {
	services := make([]string, 0, len(serviceMap))

	for service := range serviceMap {
		if service != "" {
			services = append(services, service)
		}
	}

	sort.Strings(services)

	return services
}

func mergeServiceSlices(primary []string, secondary []string) []string {
	services := primary

	for _, secService := range secondary {
		add := true
		for _, priService := range primary {
			if priService == secService {
				add = false

				break
			}
		}

		if add {
			services = append(services, secService)
		}
	}

	return services
}
