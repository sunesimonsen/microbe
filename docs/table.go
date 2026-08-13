package docs

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

var TablePage = NewPage(
	"Table",
	NewExample(
		"Default",
		`
    <table>
      <thead>
        <tr>
          <th scope="col">Planet</th>
          <th scope="col">Diameter (km)</th>
          <th scope="col">Distance to Sun (AU)</th>
          <th scope="col">Orbit (days)</th>
        </tr>
      </thead>
      <tbody>
        <tr>
          <th scope="row">Mercury</th>
          <td>4,880</td>
          <td>0.39</td>
          <td>88</td>
        </tr>
        <tr>
          <th scope="row">Venus</th>
          <td>12,104</td>
          <td>0.72</td>
          <td>225</td>
        </tr>
        <tr>
          <th scope="row">Earth</th>
          <td>12,742</td>
          <td>1.00</td>
          <td>365</td>
        </tr>
        <tr>
          <th scope="row">Mars</th>
          <td>6,779</td>
          <td>1.52</td>
          <td>687</td>
        </tr>
      </tbody>
      <tfoot>
        <th scope="row">Average</th>
        <td>9,126</td>
        <td>0.91</td>
        <td>341</td>
      </tfoot>
    </table>
    `,
	),
	NewExample(
		"Striped",
		`
    <table class="striped">
      <thead>
        <tr>
          <th scope="col">Planet</th>
          <th scope="col">Diameter (km)</th>
          <th scope="col">Distance to Sun (AU)</th>
          <th scope="col">Orbit (days)</th>
        </tr>
      </thead>
      <tbody>
        <tr>
          <th scope="row">Mercury</th>
          <td>4,880</td>
          <td>0.39</td>
          <td>88</td>
        </tr>
        <tr>
          <th scope="row">Venus</th>
          <td>12,104</td>
          <td>0.72</td>
          <td>225</td>
        </tr>
        <tr>
          <th scope="row">Earth</th>
          <td>12,742</td>
          <td>1.00</td>
          <td>365</td>
        </tr>
        <tr>
          <th scope="row">Mars</th>
          <td>6,779</td>
          <td>1.52</td>
          <td>687</td>
        </tr>
      </tbody>
      <tfoot>
        <th scope="row">Average</th>
        <td>9,126</td>
        <td>0.91</td>
        <td>341</td>
      </tfoot>
    </table>
    `,
	).WithDescription(
		P(Text("Add the "), InlineCodeList("striped"), Text(" class to color every second row to make the table more readable.")),
	),
).WithDescription(
	P(Text("Tables let you present tabular data organized into rows and columns, making it easy for users to compare related values at a glance.")),
)
