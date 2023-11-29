
console.time("tempoEsecuzione");

//first part
let input = "hepxcrrq";
//second part
//let input = "hepxxyzz";

input = input.split("");
let password = nextPassword(input);
let counter = 0;
while(!checkValidPassword(password) && counter < 1000000){
    password = nextPassword(password);
    counter++;
}

console.log(counter);
password = password.join("")
console.log(password);

console.timeEnd("tempoEsecuzione");


function checkValidPassword(password){
    if (checkTwoDoubleLetters(password) && checkValidCharacters(password) && checkValidSequence(password)){
        return true;
    }
    return false;
}


function checkValidCharacters(password){
    if (password.includes("i") || password.includes("o") || password.includes("l")){
        return false;
    }
    return true;
}


function nextPassword(password){

    for(let counter = password.length -1; counter >=0 ; counter--){
        let letter = password[counter];
        letter = getNextLetter(letter);
        password[counter] = letter;
        if (password[counter] != 'a'){
            return password;
        }
    }

    return password
}

function checkValidSequence(password){
    for(let counter = 0; counter < password.length -2; counter++){
        if (password[counter + 1].charCodeAt(0) - password[counter].charCodeAt(0) == 1
        && password[counter + 2].charCodeAt(0) - password[counter].charCodeAt(0) == 2){
            return true
        }
    }
    return false;       
}

function getNextLetter(letter){
    if (letter=='z'){
        return 'a';
    }
    let code = letter.charCodeAt(0);
    return String.fromCharCode(code + 1);
}

function checkDoubleLetters(password, excludedLetter){
    for(let counter = 0; counter < password.length -1; counter++){
        let letter = password[counter];
        if (letter == excludedLetter){
            continue;
        }
        let nextLetter = password[counter+1];
        if (letter == nextLetter){
            return letter
        }
    }

    return false
}

function checkTwoDoubleLetters(password){
    let doubleLetter = "";
    let excludedLetter = false;
    excludedLetter = checkDoubleLetters(password, doubleLetter);
    doubleLetter = checkDoubleLetters(password, excludedLetter);

    if(doubleLetter !== false){
        return true;
    }

    return false;

}