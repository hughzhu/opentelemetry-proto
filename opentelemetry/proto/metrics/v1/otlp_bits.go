//go:generate stringer -type=Kind

package tmp

type (
	Kind int
)

const (
	// One of the following three MUST be set. There are 3 exclusive Temporality kinds.
	MaskInstantaneous = 1 << 0
	MaskCumulative    = 1 << 1
	MaskDelta         = 1 << 2

	// One of the following two MUST be set. There are 2 exclusive Structure kinds.
	MaskGrouping = 1 << 3
	MaskAdding   = 1 << 4

	// May be set with MaskAdding.
	MaskMonotonic = 1 << 5

	// May be set for any instrument.
	MaskSynchronous = 1 << 6
)

const (
	Invalid Kind = 0

	// There are 18 valid Kind values (3 * 2 * 2 * 1.5), where 1.5
	// accounts for Monotonic only applying to Adding instruments

	AddingMonotonicInstantaneousSynchronous  Kind = MaskAdding | MaskInstantaneous | MaskMonotonic | MaskSynchronous
	AddingMonotonicInstantaneousAsynchronous Kind = MaskAdding | MaskInstantaneous | MaskMonotonic
	AddingMonotonicCumulativeSynchronous     Kind = MaskAdding | MaskCumulative | MaskMonotonic | MaskSynchronous
	AddingMonotonicCumulativeAsynchronous    Kind = MaskAdding | MaskCumulative | MaskMonotonic
	AddingMonotonicDeltaSynchronous          Kind = MaskAdding | MaskDelta | MaskMonotonic | MaskSynchronous
	AddingMonotonicDeltaAsynchronous         Kind = MaskAdding | MaskDelta | MaskMonotonic

	AddingInstantaneousSynchronous  Kind = MaskAdding | MaskInstantaneous | MaskSynchronous
	AddingInstantaneousAsynchronous Kind = MaskAdding | MaskInstantaneous
	AddingCumulativeSynchronous     Kind = MaskAdding | MaskCumulative | MaskSynchronous
	AddingCumulativeAsynchronous    Kind = MaskAdding | MaskCumulative
	AddingDeltaSynchronous          Kind = MaskAdding | MaskDelta | MaskSynchronous
	AddingDeltaAsynchronous         Kind = MaskAdding | MaskDelta

	GroupingInstantaneousSynchronous  Kind = MaskGrouping | MaskInstantaneous | MaskSynchronous
	GroupingInstantaneousAsynchronous Kind = MaskGrouping | MaskInstantaneous
	GroupingCumulativeSynchronous     Kind = MaskGrouping | MaskCumulative | MaskSynchronous
	GroupingCumulativeAsynchronous    Kind = MaskGrouping | MaskCumulative
	GroupingDeltaSynchronous          Kind = MaskGrouping | MaskDelta | MaskSynchronous
	GroupingDeltaAsynchronous         Kind = MaskGrouping | MaskDelta
)

// The 18 values above were generated by:
//
// var (
// 	addingProps     = []string{"Adding", "Grouping"}
// 	cumulativeProps = []string{"Instantaneous", "Cumulative", "Delta"}
// 	monoProps       = []string{"Monotonic", ""}
// 	syncProps       = []string{"Synchronous", "Asynchronous"}
// )
//
// func main() {
// 	for _, a := range addingProps {
// 		for _, m := range monoProps {
// 			if a == "Grouping" && m == "Monotonic" {
// 				continue
// 			}
// 			mfrag := ""
// 			if m == "Monotonic" {
// 				mfrag = fmt.Sprint(" | Mask", m)
// 			}
// 			for _, c := range cumulativeProps {
// 				for _, s := range syncProps {
// 					sfrag := ""
// 					if s == "Synchronous" {
// 						sfrag = fmt.Sprint(" | Mask", s)
// 					}
// 					fmt.Print(a, m, c, s, " Kind = Mask", a, " | Mask", c, mfrag, sfrag, "\n")
// 				}
// 			}
// 		}
// 	}
// }