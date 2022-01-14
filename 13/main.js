let myArray = [];
//fuction sarng object kep khr moun nuk suek sa
function info(id, name, English, Math, Physic) {
  this.id = id;
  this.name = name;
  this.score = {
    English,
    Math,
    Physic,
  };
}
//fuction srk ha khar avg score
function avgScore(obj) {
  var avg = 0;
  for (let x in obj) {
    avg += obj[x];
  }
  avg = Math.round(avg / Object.keys(obj).length);
  return avg;
}
//fuction tut grade
function grade(score) {
  let avg = avgScore(score);
  if (avg >= 80 && avg <= 100) {
    return "A";
  } else if (avg >= 70 && avg <= 79) {
    return "B";
  } else if (avg >= 60 && avg <= 69) {
    return "C";
  } else if (avg >= 50 && avg <= 59) {
    return "D";
  } else if (avg < 50) {
    return "F";
  } else {
    return "Not found!!";
  }
}
//fuction sum lup teacher user
function TeacherUser() {
  let ID = prompt(`Input student ID: `);
  let sName = prompt(`Input Name and Surename student:`);
  let english = prompt(`input score of English:`);
  let math = prompt(`input score of Math:`);
  let physic = prompt(`input score of Physic:`);
  let result = new info(
    ID,
    sName,
    parseInt(english),
    parseInt(math),
    parseInt(physic)
  );
  myArray.push(result);
}
//fucrion sum lup student user
function StudentUser() {
  let fromStudent = prompt(`Input your ID student:`);
  for (let i = 0; i < myArray.length; i++) {
    if (fromStudent === myArray[i].id) {
      console.log("===========================================");
      const { id, name, score } = myArray[i];
      console.log(`ID = ${id}`);
      console.log(`Name and surename = ${name}`);
      for (let key in score) {
        console.log(`${key} = ${score[key]}`);
      }
      console.log(`Grade = ${grade(score)}`);
      console.log("===========================================");
    } else if (fromStudent !== myArray[i].id) {
      continue;
    }
  }
}
