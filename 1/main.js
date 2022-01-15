(() => {
    let x = "myyyyyyyloooovvvveeeeeee";
    var newX = [];
    for (let i = 0; i < x.length; i++) {
        if(x[i] == x[i+1]) {
            continue;
        }
        // newX[i] = x[i];
        newX.push(x[i])
    }
    let result = newX.join("");
    console.log(result);
})();