package calculatedurationdate
import ("testing")
func Test_TransformMinutesToSeconds_Input_21718080_Should_Be_1303084800 (t *testing.T){
inputMinutes := int64(21718080)
expectSeconds := int64(1303084800)
actualSeconds := TransformMinutesToSeconds(inputMinutes)
if expectSeconds != actualSeconds {
	t.Errorf("Should be %d but actual seconds %d",expectSeconds,actualSeconds)
   }
}

func Test_TransformMinutesToSeconds_Input_15462720_Should_Be_927763200 (t *testing.T){
	inputMinutes := int64(15462720)
	expectSeconds := int64(927763200)
	actualSeconds := TransformMinutesToSeconds(inputMinutes)
	if expectSeconds != actualSeconds {
		t.Errorf("Should be %d but actual seconds %d",expectSeconds,actualSeconds)
	   }
	}

func Test_TransformMinutesToSeconds_Input_16745760_Should_Be_1004745600 (t *testing.T){
	inputMinutes := int64(16745760)
	expectSeconds := int64(1004745600)
	actualSeconds := TransformMinutesToSeconds(inputMinutes)
	if expectSeconds != actualSeconds {
		t.Errorf("Should be %d but actual seconds %d",expectSeconds,actualSeconds)
		}
	}