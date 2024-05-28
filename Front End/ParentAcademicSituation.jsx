import { useOutletContext } from "react-router-dom";
import StudentGrades from "./StudentGrades";

export default function ParentAcademicSituation() {
  let context = useOutletContext();
  console.log(context);
  let catalog = context["catalog"];
  let note = catalog["Note"];
  let role = context["rol"];
  console.log(note);
  return (
    <>
      <h2>Student's academic situation</h2>
      <StudentGrades note={note} clasa={role["copil"]["clasa"]} />
    </>
  );
}
