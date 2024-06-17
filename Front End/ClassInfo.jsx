import { Form, useLoaderData } from "react-router-dom";
import { professorLoader } from "./Professor";
import Plot from "react-plotly.js";

async function getStudents(id_clasa, id_scoala) {
  let formData = new FormData();
  formData.append("id_scoala", id_scoala);
  formData.append("id_clasa", id_clasa);
  const url = "/api/eleviClasa";
  let students = [];
  await fetch(url, {
    method: "POST",
    body: formData,
  })
    .then((response) => response.json())
    .then((data) => (students = data))
    .catch((error) => console.log(error));

  return students;
}

export async function loader({ params }) {
  const data = await professorLoader({ params });
  console.log(data);

  const classes = data["classes"];
  const role = data["role"];
  let thisClass = null;

  for (let obj of classes) if (obj["clasa"] === params.idClass) thisClass = obj;

  if (thisClass === null) {
    console.log("Throwing");
    throw new Response("Not Found", { status: 404 });
  }
  const students = await getStudents(thisClass["clasa"], role["id"]);

  let plots = [];

  let plotBoxNoteActivitate = null;
  const url =
    "/api/getNoteActivitate?id_scoala=" +
    role["id"] +
    "&id_clasa=" +
    thisClass["clasa"] +
    "&materie=" +
    thisClass["disciplina"];
  await fetch(url)
    .then((response) => response.json())
    .then(
      (data) =>
        (plotBoxNoteActivitate = <Plot data={data.data} layout={data.layout} />)
    )
    .catch((error) => console.log(error));
  plots.push(plotBoxNoteActivitate);

  let feedbackPoints = null;
  const url2 =
    "/api/GetFeedbackPoints?id_scoala=" +
    role["id"] +
    "&id_clasa=" +
    thisClass["clasa"] +
    "&materie=" +
    thisClass["disciplina"];
  await fetch(url2)
    .then((response) => response.json())
    .then(
      (data) =>
        (feedbackPoints = <Plot data={data.data} layout={data.layout} />)
    )
    .catch((error) => console.log(error));
  plots.push(feedbackPoints);

  let pieFeedback = null;
  const url3 =
    "/api/GetPieFeedback?id_scoala=" +
    role["id"] +
    "&id_clasa=" +
    thisClass["clasa"] +
    "&materie=" +
    thisClass["disciplina"];
  await fetch(url3)
    .then((response) => response.json())
    .then(
      (data) => (pieFeedback = <Plot data={data.data} layout={data.layout} />)
    )
    .catch((error) => console.log(error));
  plots.push(pieFeedback);

  let evolutie = null;
  const date = new Date("2023-10-01");
  const formatter = new Intl.DateTimeFormat("en-US", {
    year: "numeric",
    month: "2-digit",
    day: "2-digit",
  });
  const [{ value: month }, , { value: day }, , { value: year }] =
    formatter.formatToParts(date);
  const formattedDate = `${year}-${month}-${day}`;
  const url4 =
    "/api/getEvolutie?id_scoala=" +
    role["id"] +
    "&id_clasa=" +
    thisClass["clasa"] +
    "&materie=" +
    thisClass["disciplina"] +
    "&data=" +
    formattedDate;

  await fetch(url4)
    .then((response) => response.json())
    .then((data) => (evolutie = <Plot data={data.data} layout={data.layout} />))
    .catch((error) => console.log(error));
  plots.push(evolutie);

  let evolutieNoteElevi = null;
  const url5 =
    "/api/GetEvolNoteElevi?id_scoala=" +
    role["id"] +
    "&id_clasa=" +
    thisClass.clasa;
  await fetch(url5)
    .then((response) => response.json())
    .then(
      (data) =>
        (evolutieNoteElevi = <Plot data={data.data} layout={data.layout} />)
    )
    .catch((error) => console.log(error));
  plots.push(evolutieNoteElevi);

  return {
    thisClass: thisClass,
    role: role,
    students: students,
    plots: plots,
  };
}

export async function action({ request }) {
  let formData = await request.formData();
  let intent = formData.get("intent");
  for (let pair of formData.entries()) console.log(pair[0] + ": " + pair[1]);
  console.log("hello from action");
  if (intent === "feedback") {
    const url = "/api/addFeedbackProfElev";
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

  if (intent === "grade") {
    const url = "/api/adaugareNota";
    await fetch(url, {
      method: "POST",
      body: formData,
    })
      .then((response) => {
        if (response.status === 200) alert("Added grade!");
        else alert("Error!");
        return response.json();
      })
      .then((response) => console.log(response))
      .catch((error) => console.log(error));
    return 0;
  }
  console.log("hello from action throw :(");
  throw json({ message: "Invalid intent" }, { status: 400 });
}

function groupByStudent(students) {
  console.log(students);
  return students.reduce((acc, student) => {
    const { nume, prenume, disciplina, nota, id } = student;
    const studentName = `${nume} ${prenume}`;

    if (!acc[studentName]) {
      acc[studentName] = {
        name: studentName,
        id: id,
        grades: [],
      };
    }

    acc[studentName].grades.push({ disciplina, nota });

    return acc;
  }, {});
}

export default function ClassInfo() {
  const data = useLoaderData();
  const role = data["role"];
  const thisClass = data["thisClass"];
  console.log(data);
  const groupedStudents = groupByStudent(data["students"]["data"]);

  console.log(groupedStudents);
  return (
    <>
      <h2>Class {data["thisClass"]["clasa"]}</h2>
      <h2>{data["thisClass"]["disciplina"]}</h2>
      <Form method="post">
        <h4>Grade a student</h4>
        <input type="hidden" name="id_scoala" value={role["id"]} required />
        <input
          type="hidden"
          name="id_clasa"
          value={thisClass["clasa"]}
          required
        />
        <input
          type="hidden"
          name="nume_disciplina"
          value={thisClass["disciplina"]}
          required
        />
        <label>
          Student
          <select name="id_elev" required>
            {Object.values(groupedStudents).map((stud) => (
              <option key={stud.id} value={stud.id}>
                {stud.name}
              </option>
            ))}
          </select>
        </label>
        <br />
        <label>
          Grade:{" "}
          <input type="number" min={1} max={10} name="valoare" required />
        </label>
        <br />
        <button type="submit" name="intent" value="grade">
          Submit
        </button>
      </Form>
      <br />

      <Form method="post">
        <h4>Feedback for student</h4>
        <input type="hidden" name="id_scoala" value={role["id"]} required />
        <input
          type="hidden"
          name="id_clasa"
          value={thisClass["clasa"]}
          required
        />
        <input
          type="hidden"
          name="materie"
          value={thisClass["disciplina"]}
          required
        />
        <label>
          Student
          <select name="id_elev" required>
            {Object.values(groupedStudents).map((stud) => (
              <option key={stud.id} value={stud.id}>
                {stud.name}
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
        <button type="submit" name="intent" value="feedback">
          Submit
        </button>
      </Form>

      {data.plots.map((plot, i) => {
        return <div key={i}>{plot}</div>;
      })}

      {Object.values(groupedStudents).map((student, index) => (
        <div key={index} className="student">
          <b>{student.name}</b>
          <ul style={{ listStyleType: "none" }}>
            {student.grades.map(
              (grade, idx) =>
                grade.disciplina === data.thisClass["disciplina"] && (
                  <li key={idx}>
                    {grade.disciplina}: {grade.nota}
                  </li>
                )
            )}
          </ul>
          <br />
        </div>
      ))}
    </>
  );
}
