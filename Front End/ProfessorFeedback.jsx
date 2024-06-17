import { useLoaderData } from "react-router-dom";
import { professorLoader } from "./Professor";

export async function loader({ params }) {
  const data = await professorLoader({ params });
  const role = data["role"];
  let classes = data["classes"];
  for (let i = 0; i < classes.length; i++) {
    let feedback;
    const url =
      "/api/getFeedbackProf?id_scoala=" +
      role["id"] +
      "&id_clasa=" +
      classes[i].clasa +
      "&materie=" +
      classes[i].disciplina;
    await fetch(url)
      .then((response) => response.json())
      .then((data) => (feedback = data))
      .catch((error) => console.log(error));
    classes[i]["feedback"] = feedback;
  }
  return { role: role, classes: classes };
}

export default function ProfessorFeedback() {
  const data = useLoaderData();
  console.log(data);
  const role = data["role"];
  const classes = data["classes"];
  return (
    <>
      <h2>Professor's Feedback</h2>
      {classes.map((obj) => (
        <div key={obj.clasa + obj.disciplina}>
          {obj.feedback.length>0 && (
            <>
              <p>
                Class {obj.clasa} - {obj.disciplina}
              </p>
              {obj.feedback.map((feed) => {
                return <p>{feed.tip ? "🙂" : "🙁"} {feed.content}({feed.data})</p>;
              })}
            </>
          )}
          <br />
          <br />
        </div>
      ))}
    </>
  );
}
