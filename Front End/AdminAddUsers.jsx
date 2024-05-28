import { Form, useLoaderData, useOutletContext } from "react-router-dom";

export async function action({ request, params }) {
  let formData = await request.formData();
  const id_scoala = formData.get("id_scoala");
  let formData1 = new FormData();
  formData1.append("id_scoala", id_scoala);
  let formData2 = new FormData();
  formData2.append("id_scoala", id_scoala);
  let formData3 = new FormData();
  formData3.append("id_scoala", id_scoala);
  let formData4 = new FormData();
  formData4.append("id_scoala", id_scoala);
  let formData5 = new FormData();
  formData5.append("id_scoala", id_scoala);

  for (const pair of formData.entries()) {
    console.log("Fisierul " + pair[0]);
    // console.log(typeof pair[1]);
    if (pair[0] === "csv_file1") {
      // console.log("Fisierul " + pair[0]);
      formData1.append("csv_file", pair[1]);
      const url = "/api/inserareClasa";
      await fetch(url, {
        method: "POST",
        body: formData1,
      })
        .then((response) => response.json())
        .then((response) => console.log(response))
        .catch((error) => console.log(error));
    } else if (pair[0] === "csv_file2") {
      formData2.append("csv_file", pair[1]);
      const url = "/api/inserareElev";
      await fetch(url, {
        method: "POST",
        body: formData2,
      })
        .then((response) => response.json())
        .then((response) => console.log(response))
        .catch((error) => console.log(error));
    } else if (pair[0] === "csv_file3") {
      formData3.append("csv_file", pair[1]);
      const url = "/api/inserareProfesor";
      await fetch(url, {
        method: "POST",
        body: formData3,
      })
        .then((response) => response.json())
        .then((response) => console.log(response))
        .catch((error) => console.log(error));
    } else if (pair[0] === "csv_file4") {
      formData4.append("csv_file", pair[1]);
      const url = "/api/insertIncadrare";
      await fetch(url, {
        method: "POST",
        body: formData4,
      })
        .then((response) => response.json())
        .then((response) => console.log(response))
        .catch((error) => console.log(error));
    } else if (pair[0] === "csv_file5") {
      formData5.append("csv_file", pair[1]);
      const url = "/api/insertMaterie";
      await fetch(url, {
        method: "POST",
        body: formData5,
      })
        .then((response) => response.json())
        .then((response) => console.log(response))
        .catch((error) => console.log(error));
    } else {
      console.log(pair[0]);
    }
  }
  return 0;
}

export default function AdminAddUsers() {
  const data = useOutletContext();
  console.log(data);
  return (
    <>
      <h3>Add information about the school</h3>
      <Form method="post" encType={"multipart/form-data"}>
        <label>
          Classes
          <input type="file" name="csv_file1" required />
        </label>
        <br />
        <label>
          Students
          <input type="file" accept=".csv" name="csv_file2" required />
        </label>
        <br />
        <label>
          Professors
          <input type="file" accept=".csv" name="csv_file3" required />
        </label>
        <br />
        <label>
          Subjects
          <input type="file" accept=".csv" name="csv_file5" required />
        </label>
        <br />

        <label>
          Placement
          <input type="file" accept=".csv" name="csv_file4" required />
        </label>
        <br />
        <input type="hidden" name="id_scoala" value={data["id"]} />
        <button type="submit">Submit</button>
      </Form>
    </>
  );
}
