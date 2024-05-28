export default function StudentGrades({ note, clasa }) {
  console.log(note);
  let tabelNote = [];
  if (Object.keys(note).length > 0)
    for (let [discipline, array] of Object.entries(note)) {
      for (let index = 0; index < array.length; index++) {
        tabelNote.push(
          <tr key={index}>
            <td>{discipline}</td>
            <td>{array[index]["nota"]}</td>
            <td>{array[index]["data"]}</td>
          </tr>
        );
      }
    }
  else {
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
            <th>Grade</th>
            <th>Date</th>
          </tr>
        </thead>
        <tbody>{tabelNote}</tbody>
      </table>
    </>
  );
}
