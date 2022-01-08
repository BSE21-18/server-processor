//https://github.com/cdipaolo/goml/blob/master/linear/softmax.go

//https://en.wikipedia.org/wiki/Logistic_regression
/*
 Softmax represents a softmax classification model
 in 'k' demensions. It is generally thought of as
 a generalization of the Logistic Regression model.
 
 Prediction will return a vector ([]float64) that
 corresponds to the probabilty (where i is the index)
 that the inputted features is 'i'. 
 
 Softmax classification
 operates assuming the Multinomial probablility
 distribution of data.
 */
 
 type Softmax struct {
	// alpha and maxIterations are used only for
	// GradientAscent during learning. If maxIterations
	// is 0, then GradientAscent will run until the
	// algorithm detects convergance.
	//
	// regularization is used as the regularization
	// term to avoid overfitting within regression.
	// Having a regularization term of 0 is like having
	// _no_ data regularization. The higher the term,
	// the greater the bias on the regression
	alpha          float64
	regularization float64
	maxIterations  int

	// k is the dimension of classification (the number
	// of possible outcomes)
	k int

	// method is the optimization method used when training
	// the model
	method base.OptimizationMethod

	// trainingSet and expectedResults are the
	// 'x', and 'y' of the data, expressed as
	// vectors, that the model can optimize from
	trainingSet     [][]float64
	expectedResults []float64

	Parameters [][]float64 `json:"theta"`

	// Output is the io.Writer used for logging
	// and printing. Defaults to os.Stdout.
	Output io.Writer
}









