import { Link, useLoaderData } from "react-router-dom";
import { professorLoader } from "./Professor";

export async function loader(props) {
  return professorLoader(props);
}

export default function SelectClass() {
  const data = useLoaderData();
  const role = data["role"];
  const classes = data["classes"];

  return (
    <>
      <h2>Select Class</h2>
      {classes.map((obj) => (
        <div key={obj.clasa}>
          <h3>Class {obj.clasa}</h3>
          <p>Discipline: {obj.disciplina}</p>
          <Link to={"classinfo/"+obj.clasa}><p>See students</p></Link>
          <br></br>
        </div>
      ))}
    </>
  );
}
