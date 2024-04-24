import { Form } from "react-router-dom";

export async function action({ request, params }) {
  const formData = await request.formData();
  const formData1 = new FormData();
  formData1.append("id_scoala", 1);
  const formData2 = new FormData();
  formData2.append("id_scoala", params.idAdmin);
  const formData3 = new FormData();
  formData3.append("id_scoala", params.idAdmin);
  //   const formData4 = new FormData();
  //   formData4.append("id_scoala", params.idAdmin);

  for (const pair of formData.entries()) {
    if (pair[0] === "csv_file1") {
      formData1.append("csv_file", pair[1]);
      const url = "/api/inserareClasa";
      await fetch(url, {
        method: "POST",
        body: formData1,
      })
        .then((response) => console.log(response))
        .catch((error) => console.log(error));
    } else if (pair[0] === "csv_file2") {
      formData2.append("csv_file", pair[1]);
      const url = "/api/inserareProfesor";
      await fetch(url, {
        method: "POST",
        body: formData2,
      })
        .then((response) => console.log(response))
        .catch((error) => console.log(error));
    } else if (pair[0] === "csv_file3") {
      formData3.append("csv_file", pair[1]);
      const url = "/api/inserareElev";
      await fetch(url, {
        method: "POST",
        body: formData3,
      })
        .then((response) => console.log(response))
        .catch((error) => console.log(error));
    } else {
      console.log(pair[0]);
    }
  }
  return 0;
}

export default function AdminAddUsers() {
  return (
    <>
      <h3>Add information about the school</h3>
      <Form method="post">
        <label>
          Classes
          <input type="file" accept=".csv" name="csv_file1" required />
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
        {/* <label><input type="file" accept=".csv" name="csv_file4" required/></label> */}
        <button type="submit">Submit</button>
      </Form>
    </>
  );
}
