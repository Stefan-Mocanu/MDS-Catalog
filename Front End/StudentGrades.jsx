export default function StudentGrades({ catalog, clasa }) {
  const note = catalog["Note"];
  const activitate = catalog["Activitate"];
  const absente = catalog["Absente"];
  const feedback = catalog["Feedback"];
  console.log(note);
  let tabelNote = [];
  if (
    Object.keys(note).length > 0 ||
    absente.length > 0 ||
    Object.keys(activitate).length > 0 ||
    Object.keys(feedback).length > 0
  ) {
    for (let [discipline, array] of Object.entries(note)) {
      for (let index = 0; index < array.length; index++) {
        tabelNote.push(
          <tr key={discipline + index + "grade"}>
            <td>{discipline}</td>
            <td>{array[index]["data"]}</td>
            <td>Grade: {array[index]["nota"]}</td>
          </tr>
        );
      }
    }
    for (let [discipline, array] of Object.entries(activitate)) {
      for (let index = 0; index < array.length; index++) {
        tabelNote.push(
          <tr key={discipline + index + "activity"}>
            <td>{discipline}</td>
            <td>{array[index]["data"]}</td>
            <td>Activity: {array[index]["nota"]}</td>
          </tr>
        );
      }
    }
    for (let i in absente) {
      tabelNote.push(
        <tr key={absente[i].materie + i + "absence"}>
          <td>{absente[i].materie}</td>
          <td>{absente[i].data}</td>
          <td>Absence</td>
        </tr>
      );
    }
    for (let [discipline, array] of Object.entries(feedback)) {
      for (let index = 0; index < array.length; index++) {
        tabelNote.push(
          <tr key={discipline + index + "feedback"}>
            <td>{discipline}</td>
            <td>{array[index]["data"]}</td>
            <td>Feedback {array[index].tip ? "🙂" : "🙁"}: {array[index]["content"]}</td>
          </tr>
        );
      }
    }
  } else {
    tabelNote = (
      <tr>
        <td>
          <i>Nothing</i>
        </td>
        <td>
          <i>Nothing</i>
        </td>
        <td>
          <i>Nothing</i>
        </td>
      </tr>
    );
  }

  return (
    <>
      <h3>Grades</h3>
      <h3>Class {clasa}</h3>
      <table border={1}>
        <thead>
          <tr>
            <th>Subject</th>
            <th>Date</th>
            <th>Grade/Activity</th>
          </tr>
        </thead>
        <tbody>{tabelNote}</tbody>
      </table>
    </>
  );
}
