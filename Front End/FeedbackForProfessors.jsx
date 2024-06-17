import { Form, useLoaderData } from "react-router-dom";
import { studentLoader } from "./Student";

export async function loader({ params }) {
  const data = await studentLoader({ params });
  return data;
}

export async function action({ request }) {
  const formData = await request.formData();
  const url = "/api/addFeedbackElevProf";
  await fetch(url, {
    method: "POST",
    body: formData,
  })
    .then((response) => {
      if (response.status === 200) alert("Feedback sent!");
      else alert("Error!");
      return response.json();
    })
    .then((response) => console.log(response))
    .catch((error) => console.log(error));
  return 0;
}

export default function FeedbackForProfessors() {
  const data = useLoaderData();
  const role = data["rol"];
  const materii = [];
  for (let i in data.cataloage)
    materii.push(...Object.keys(data.cataloage[i].catalog.Note));
  console.log(data);
  return (
    <>
      <h2>Feedback for professors</h2>
      <Form method="post">
        <input type="hidden" name="id_scoala" value={role["id"]} required />
        <label>
          Class
          <select name="id_clasa" required>
            {data.cataloage.map((obj) => {
              return (
                <option key={obj.id_clasa} value={obj.id_clasa}>
                  {obj.id_clasa}
                </option>
              );
            })}
          </select>
        </label>
        <br />
        <label>
          Discipline
          <select name="materie" required>
            {materii.map((materie) => (
              <option key={materie} value={materie}>
                {materie}
              </option>
            ))}
          </select>
        </label>
        <br />
        <label>
          🙂
          <input type="radio" name="tip" value="1" defaultChecked />
        </label>
        <br />
        <label>
          🙁
          <input type="radio" name="tip" value="0" />
        </label>
        <br />
        <label>
          Feedback:
          <br />
          <textarea name="continut" rows="5" required />
        </label>
        <br />
        <button type="submit">Submit</button>
      </Form>
    </>
  );
}
