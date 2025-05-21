costFunction = document.getElementById("costFunction")
costFunction2 = document.getElementById("costFunction2")
wbParameters = document.getElementById("wbParameters")
modelPrediction = document.getElementById("modelPrediction")
actualResult = document.getElementById("actualResult")
wibiParameters = document.getElementById("wibiParameters")
costFunctionDefinition = document.getElementById("costFunctionDefinition")
modelPredictionDefinition = document.getElementById("modelPredictionDefinition")
gradientDescentDefinition = document.getElementById("gradientDescentDefinition")
partialDerivative = document.getElementById("partialDerivative")

katex.render("\\textbf{J}(\\vec w,b)=", costFunction, {
    throwOnError: false
});

katex.render("\\textbf{J}(\\vec w,b)", costFunction2, {
    throwOnError: false
});

katex.render("\\vec w,b",wbParameters,{
    throwOnError: false
});

katex.render("f\\vec w,b(\\vec x)",modelPrediction,{
    throwOnError: false
});

katex.render("y^{(i)}",actualResult,{
    throwOnError: false
});

katex.render("x^{(i)},y^{(i)}",wibiParameters,{
    throwOnError: false
});

katex.render("J(\\vec w,b) = \\frac{1 }{2m}\\sum\\limits_{i = 1}^{m}(f\\vec w,b(\\vec x)-y^{(i)})^{2}",costFunctionDefinition,{
    throwOnError: false
});

katex.render("f\\vec w,b(\\vec x) = \\vec w \\cdot \\vec x + b ",modelPredictionDefinition,{
    throwOnError: false
});

katex.render(String.raw`
 \text{repeat until convergence: }  \lbrace \newline 
          \hspace{1cm} wj = wj - \alpha \frac{\partial J(w,b)}{\partial wj} \hspace{10mm} \small{j=1,2...n}\\
          \hspace{1cm} b = b - \alpha \frac{\partial J(w,b)}{\partial b} \newline
          \rbrace`,gradientDescentDefinition,{
    throwOnError: false
});

katex.render(String.raw`
 \frac{\partial J(\mathbf{w},b)}{\partial w_j}  = \frac{1}{m} \sum\limits_{i = 1}^{m} (f_{\mathbf{w},b}(\mathbf{x}^{(i)}) - y^{(i)})x_{j}^{(i)} \\
 \frac{\partial J(\mathbf{w},b)}{\partial b}  = \frac{1}{m} \sum\limits_{i = 1}^{m} (f_{\mathbf{w},b}(\mathbf{x}^{(i)}) - y^{(i)})
 `,partialDerivative,{
    throwOnError: false
});