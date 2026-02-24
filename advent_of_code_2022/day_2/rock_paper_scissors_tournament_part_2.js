import data from "./data.js";
import {
  getShapeByLetter,
  myMatchResult,
  getMatchResultPoints,
  getMatchShapePoints,
} from "./rock_paper_scissors_tournament_part_1.js";
var result = 0;

// What should this match result be
// X => Lose
// Y => Draw
// Z => Win
export const getMatchOutcomes = (letter) => {
  switch (letter) {
    case "X":
      return "Lose";

    case "Y":
      return "Draw";

    case "Z":
      return "Win";
  }
};
const getMyShapeBasedbyOutcomes = (outcome, oppenentShape) => {
  switch (oppenentShape) {
    case "Rock":
      return outcome === "Win"
        ? "Paper"
        : outcome === "Draw"
        ? "Rock"
        : "Scissors";

    case "Paper":
      return outcome === "Win"
        ? "Scissors"
        : outcome === "Draw"
        ? "Paper"
        : "Rock";

    case "Scissors":
      return outcome === "Win"
        ? "Rock"
        : outcome === "Draw"
        ? "Scissors"
        : "Paper";
  }
};

// Get current Match winner
// How to cacluate score for each round
// The shape + the match outcome
const matchLogic = (match) => {
  var [oppenent, outcome] = match;

  oppenent = getShapeByLetter(oppenent);
  outcome = getMatchOutcomes(outcome);
  var myShape = getMyShapeBasedbyOutcomes(outcome, oppenent);

  const matchResult = myMatchResult(myShape, oppenent);
  const matchScore =
    getMatchResultPoints(matchResult) + getMatchShapePoints(myShape);

  console.table({
    oppenent,
    outcome,
    myShape,
  });
  return matchScore;
};

for (var i = 0; i <= data.length; i++) {
  if (Array.isArray(data[i])) result = matchLogic(data[i]) + result;
}


console.log(result); // 14610
