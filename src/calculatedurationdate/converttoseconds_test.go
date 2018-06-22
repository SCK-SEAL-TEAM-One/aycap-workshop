package calculatedurationdate
import ("testing")
func Test_TransformMinutesToSeconds_Input_21718080_Should_Be_1303084800 (t *testing.T){
inputMinutes := 21718080
expectSeconds := 1303084800
actualSeconds := TransformMinutesToSeconds(inputMinutes)
if expectSeconds != actualSeconds {
	t.Errorf("Should be %d but actual seconds %d",expectSeconds,actualSeconds)
   }
}

func Test_TransformMinutesToSeconds_Input_15462720_Should_Be_927763200 (t *testing.T){
	inputMinutes := 15462720
	expectSeconds := 927763200
	actualSeconds := TransformMinutesToSeconds(inputMinutes)
	if expectSeconds != actualSeconds {
		t.Errorf("Should be %d but actual seconds %d",expectSeconds,actualSeconds)
	   }
	}