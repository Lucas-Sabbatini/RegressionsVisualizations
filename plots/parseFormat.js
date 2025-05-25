/**
 * Parses a string containing monomials and returns an array of monomial strings.
 * (No changes needed in this function for the request)
 */
function parseMonomialString(monomialStr) {
    if (!monomialStr || typeof monomialStr !== 'string') {
        return [];
    }
    let cleanedStr = monomialStr;
    if (monomialStr.startsWith("{") && monomialStr.endsWith("}")) {
        cleanedStr = monomialStr.substring(1, monomialStr.length - 1);
    }
    if (!cleanedStr) {
        return [];
    }
    const monomialList = cleanedStr.split(',').map(m => m.trim()).filter(m => m.length > 0);
    return monomialList;
}

/**
 * Formats a single monomial string (e.g., "x1", "x2^2", "x1x2") into LaTeX.
 * (No changes needed in this function for the request)
 */
function formatSingleMonomialToLatex(monomial) {
    if (!monomial) return "";
    let latexResult = "";
    let currentIndex = 0;
    while (currentIndex < monomial.length) {
        let variableBase = "";
        let variableSubscript = "";
        let variableExponent = "";
        if (monomial[currentIndex] && monomial[currentIndex].match(/[a-zA-Z]/)) {
            variableBase = monomial[currentIndex];
            currentIndex++;
            while (currentIndex < monomial.length &&
                   (monomial[currentIndex].match(/[0-9]/) ||
                    (variableSubscript.length === 0 && monomial[currentIndex].match(/[a-zA-Z]/) && !monomial[currentIndex+1]?.match(/\^|[0-9]/) )
                   )
                  ) {
                if (monomial[currentIndex].match(/[a-zA-Z]/) && variableSubscript.length > 0 && variableSubscript.match(/[0-9]$/)) {
                    break;
                }
                variableSubscript += monomial[currentIndex];
                currentIndex++;
            }
        } else {
            if (latexResult.length > 0) latexResult += " ";
            latexResult += monomial[currentIndex];
            currentIndex++;
            continue;
        }
        let formattedVar = "";
        let actualBase = variableBase;
        let actualSubscript = "";
        let k = 0;
        while(k < variableSubscript.length && variableSubscript[k].match(/[a-zA-Z]/)) {
            actualBase += variableSubscript[k];
            k++;
        }
        actualSubscript = variableSubscript.substring(k);
        formattedVar = actualBase;
        if (actualSubscript && actualSubscript.match(/^[0-9]+$/)) {
            formattedVar += `_{${actualSubscript}}`;
        } else if (actualSubscript) {
             formattedVar += actualSubscript;
        }
        if (currentIndex < monomial.length && monomial[currentIndex] === '^') {
            currentIndex++;
            while (currentIndex < monomial.length && monomial[currentIndex].match(/[0-9]/)) {
                variableExponent += monomial[currentIndex];
                currentIndex++;
            }
            if (variableExponent) {
                formattedVar += `^{${variableExponent}}`;
            } else {
                formattedVar += "^{}";
            }
        }
        if (latexResult.length > 0 && formattedVar.length > 0) {
            latexResult += " ";
        }
        latexResult += formattedVar;
    }
    return latexResult;
}

/**
 * Helper function to format a number to a string with up to two decimal places.
 * Removes trailing ".00" and ".X0".
 * @param {number} num - The number to format.
 * @returns {string} The formatted number as a string.
 */
function formatNumberToMaxTwoDecimals(num) {
    if (typeof num !== 'number') {
        return String(num); 
    }
    const roundedNum = num.toFixed(3);
    return String(roundedNum);
}

/**
 * Updates the prediction function display with formatted coefficients and bias.
 *
 * @param {number[]} coefficients - An array of numbers for monomial coefficients.
 * @param {number} bias - The numerical bias term.
 */
export function updatePedictionFunction(coefficients, bias) {
    const predictionFunction = document.getElementById("predictionFunction");
    if (!predictionFunction) {
        console.error("Element with ID 'predictionFunction' not found.");
        return;
    }

    console.log("Raw coefficients:", coefficients, "Raw bias:", bias);
    let input = document.getElementById("modeloLinear").value;
    let monomials = parseMonomialString(input);

    const formattedStandaloneBias = formatNumberToMaxTwoDecimals(bias);

    if (!monomials || monomials.length === 0) {
        katex.render(formattedStandaloneBias, predictionFunction, { throwOnError: false });
        return;
    }

    const latexTerms = [];

    for (let i = 0; i < monomials.length; i++) {
        const coeff = (coefficients && typeof coefficients[i] === 'number') ? coefficients[i] : 1;

        if (coeff === 0) {
            continue;
        }

        const currentMonomialString = monomials[i];
        const latexMonomial = formatSingleMonomialToLatex(currentMonomialString);

        if (!latexMonomial) continue;

        let termString = "";
        const absCoeff = Math.abs(coeff);
        
        let coeffDisplayValue;
        if (absCoeff === 1) {
            coeffDisplayValue = ""; 
        } else {
            coeffDisplayValue = formatNumberToMaxTwoDecimals(absCoeff) + " ";
        }

        let currentCoeffDisplaySign = "";
        if (coeff < 0) {
            currentCoeffDisplaySign = "-";
        }
        
        if (latexTerms.length === 0) { 
            termString = currentCoeffDisplaySign + coeffDisplayValue + latexMonomial;
        } else { 
            if (coeff < 0) {
                termString = ` - ${coeffDisplayValue}` + latexMonomial; 
            } else { 
                termString = ` + ${coeffDisplayValue}` + latexMonomial; 
            }
        }
        latexTerms.push(termString);
    }

    let polynomialString = latexTerms.join("").trim();

    if (bias !== 0) {
        const formattedAbsBias = formatNumberToMaxTwoDecimals(Math.abs(bias));
        if (polynomialString === "") {
            polynomialString = formattedStandaloneBias; 
        } else {
            if (bias < 0) {
                polynomialString += ` - ${formattedAbsBias}`;
            } else { 
                polynomialString += ` + ${formattedAbsBias}`;
            }
        }
    } else { 
        if (polynomialString === "") {
            polynomialString = "0"; 
        }
    }

    katex.render(polynomialString, predictionFunction, {
        throwOnError: false
    });
}