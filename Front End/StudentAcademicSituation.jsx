import { useOutletContext } from "react-router-dom";
import StudentGrades from "./StudentGrades";
export default function StudentAcademicSituation() {
  let context = useOutletContext();
  let cataloage = context["cataloage"];
  let role = context["rol"];

  console.log(cataloage);
  return cataloage.map((element, key) => {
    return (
      <StudentGrades
        key={key}
        catalog={element["catalog"]}
        clasa={element["id_clasa"]}
      />
    );
  });
}
