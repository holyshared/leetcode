package main

type Pair struct {
	key   string
	value float64
}

func find(gidWeight map[string]*Pair, nodeId string) *Pair {
	entry, ok := gidWeight[nodeId]
	if !ok {
		gidWeight[nodeId] = &Pair{
			key:   nodeId,
			value: 1.0,
		}
		entry = gidWeight[nodeId]
	}

	if entry.key != nodeId {
		newEntry := find(gidWeight, entry.key)
		gidWeight[nodeId] = &Pair{
			key:   newEntry.key,
			value: entry.value * newEntry.value,
		}
	}
	entry, _ = gidWeight[nodeId]

	return entry
}

func union(gidWeight map[string]*Pair, dividend string, divisor string, value float64) {
	dividendEntry := find(gidWeight, dividend)
	divisorEntry := find(gidWeight, divisor)

	dividendGid := dividendEntry.key
	divisorGid := divisorEntry.key
	dividendWeight := dividendEntry.value
	divisorWeight := divisorEntry.value

	if dividendGid != divisorGid {
		gidWeight[dividendGid] = &Pair{
			key:   divisorGid,
			value: divisorWeight * value / dividendWeight,
		}
	}
}

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	gidWeight := map[string]*Pair{}

	for i := 0; i < len(equations); i++ {
		equation := equations[i]
		dividend, divisor := equation[0], equation[1]
		quotient := values[i]

		union(gidWeight, dividend, divisor, quotient)
	}

	results := make([]float64, len(queries))

	for i := 0; i < len(queries); i++ {
		query := queries[i]
		dividend, divisor := query[0], query[1]

		_, ok1 := gidWeight[dividend]
		_, ok2 := gidWeight[divisor]

		if !ok1 || !ok2 {

			results[i] = -1.0
		} else {
			dividendEntry := find(gidWeight, dividend)
			divisorEntry := find(gidWeight, divisor)

			dividendGid := dividendEntry.key
			divisorGid := divisorEntry.key
			dividendWeight := dividendEntry.value
			divisorWeight := divisorEntry.value

			if dividendGid != divisorGid {
				results[i] = -1.0
			} else {
				results[i] = dividendWeight / divisorWeight
			}
		}
	}

	return results
}

/*
class Solution {

	public double[] calcEquation(List<List<String>> equations, double[] values,
					List<List<String>> queries) {

			HashMap<String, Pair<String, Double>> gidWeight = new HashMap<>();

			// Step 1). build the union groups
			for (int i = 0; i < equations.size(); i++) {
					List<String> equation = equations.get(i);
					String dividend = equation.get(0), divisor = equation.get(1);
					double quotient = values[i];

					union(gidWeight, dividend, divisor, quotient);
			}

			// Step 2). run the evaluation, with "lazy" updates in find() function
			double[] results = new double[queries.size()];
			for (int i = 0; i < queries.size(); i++) {
					List<String> query = queries.get(i);
					String dividend = query.get(0), divisor = query.get(1);

					if (!gidWeight.containsKey(dividend) || !gidWeight.containsKey(divisor))
							// case 1). at least one variable did not appear before
							results[i] = -1.0;
					else {
							Pair<String, Double> dividendEntry = find(gidWeight, dividend);
							Pair<String, Double> divisorEntry = find(gidWeight, divisor);

							String dividendGid = dividendEntry.getKey();
							String divisorGid = divisorEntry.getKey();
							Double dividendWeight = dividendEntry.getValue();
							Double divisorWeight = divisorEntry.getValue();

							if (!dividendGid.equals(divisorGid))
									// case 2). the variables do not belong to the same chain/group
									results[i] = -1.0;
							else
									// case 3). there is a chain/path between the variables
									results[i] = dividendWeight / divisorWeight;
					}
			}

			return results;
	}

	private Pair<String, Double> find(HashMap<String, Pair<String, Double>> gidWeight, String nodeId) {
			if (!gidWeight.containsKey(nodeId))
					gidWeight.put(nodeId, new Pair<String, Double>(nodeId, 1.0));

			Pair<String, Double> entry = gidWeight.get(nodeId);
			// found inconsistency, trigger chain update
			if (!entry.getKey().equals(nodeId)) {
					Pair<String, Double> newEntry = find(gidWeight, entry.getKey());
					gidWeight.put(nodeId, new Pair<String, Double>(
									newEntry.getKey(), entry.getValue() * newEntry.getValue()));
			}

			return gidWeight.get(nodeId);
	}

	private void union(HashMap<String, Pair<String, Double>> gidWeight, String dividend, String divisor, Double value) {
			Pair<String, Double> dividendEntry = find(gidWeight, dividend);
			Pair<String, Double> divisorEntry = find(gidWeight, divisor);

			String dividendGid = dividendEntry.getKey();
			String divisorGid = divisorEntry.getKey();
			Double dividendWeight = dividendEntry.getValue();
			Double divisorWeight = divisorEntry.getValue();

			// merge the two groups together,
			// by attaching the dividend group to the one of divisor
			if (!dividendGid.equals(divisorGid)) {
					gidWeight.put(dividendGid, new Pair<String, Double>(divisorGid,
									divisorWeight * value / dividendWeight));
			}
	}
}
*/
