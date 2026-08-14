package docs

var AccordionPage = NewPage(
	"Accordion",
	`<p>An element that organizes content into a vertically stacked list of collapsible sections. Users can click or tap a section's header to expand it and reveal detailed information, or collapse it to hide the content and reduce scrolling.</p>`,
	NewExample(
		"Example",
		`<p>Use an accordion to progressively disclose long-form content, letting users open only the sections that interest them instead of reading through everything at once.</p>`,
		`
    <p>
      Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nulla
      ullamcorper blandit ultricies. Etiam non suscipit felis. Orci varius
      natoque penatibus et magnis dis parturient montes, nascetur ridiculus
      mus. Proin feugiat purus hendrerit sapien condimentum, eu facilisis ex
      eleifend. Sed lobortis est a urna accumsan vestibulum. Curabitur iaculis
      sem lacus, id molestie eros ultrices eu. Nullam justo nulla, sollicitudin
      sed mi pretium, consequat varius erat. Phasellus sed eros dictum, congue
      enim in, dictum ante. Sed condimentum ut elit eu vehicula.
    </p>
    <div class="accordion">
      <details>
        <summary>Section 1</summary>
        <p>
          Fusce congue nec massa id eleifend. Donec id luctus ligula. In
          pellentesque diam eu magna interdum, a tristique arcu dignissim. In
          eu lacinia nisl. Phasellus eu nisi vitae enim aliquam rutrum.
          Suspendisse fringilla tortor et tincidunt vulputate. Nam a urna a
          purus ornare gravida non sit amet risus. Morbi in justo quis velit
          elementum fermentum. Etiam tristique diam nunc, quis suscipit velit
          suscipit non.
        </p>
      </details>
      <details>
        <summary>Section 2</summary>
        <p>
          Nam at mollis ante, non rutrum enim. Nulla facilisi. Nunc maximus
          diam a dui tincidunt vulputate. Phasellus in suscipit augue. In quis
          elementum enim. Mauris at porttitor libero. Sed vel tortor sit amet
          felis porttitor dapibus et id neque.
        </p>
      </details>
      <details>
        <summary>Section 3</summary>
        <p>
          Duis facilisis eros porttitor mauris sollicitudin dapibus. Mauris
          mollis turpis nec ligula pulvinar finibus. Proin molestie varius
          nisl. Ut et ex lobortis, auctor nunc tincidunt, aliquam sem. Nunc
          sodales aliquam magna. Proin rutrum sed erat accumsan elementum.
          Mauris posuere augue non orci maximus, sed sagittis elit mattis.
          Proin mattis, leo at luctus ultrices, nulla dui bibendum ipsum, sed
          venenatis ex ipsum a erat. Ut libero dui, viverra in purus eget,
          efficitur pellentesque justo.
        </p>
      </details>
      <details>
        <summary>Section 4</summary>
        <p>
          Suspendisse ultricies tristique elit eget suscipit. Sed eu odio eu
          diam convallis venenatis sit amet quis elit. Aenean mi diam, gravida
          eu nulla eget, egestas ultricies velit. Pellentesque et metus
          fringilla, lobortis felis id, suscipit nunc. Nunc interdum placerat
          dui, quis varius turpis tempor in. Maecenas faucibus tellus non lacus
          pulvinar, a malesuada urna luctus. Phasellus fermentum porta dolor
          vel hendrerit. Donec facilisis maximus ipsum congue mattis. Nullam
          justo arcu, aliquet id turpis sed, maximus laoreet dui. Nam nec elit
          accumsan, facilisis lorem ac, ullamcorper arcu. Interdum et malesuada
          fames ac ante ipsum primis in faucibus. Nullam quis rutrum libero, a
          sollicitudin magna.
        </p>
      </details>
    </div>
    <p>
      Pellentesque at sodales nisl, in eleifend erat. Pellentesque porttitor
      lacinia nibh, a dictum orci condimentum sed. Quisque eu maximus leo.
      Vestibulum elit arcu, molestie eget posuere vel, dapibus pharetra elit.
      Aenean eget turpis et turpis ullamcorper pretium. Praesent rutrum
      sollicitudin leo quis tincidunt. Aliquam erat volutpat. Ut viverra, dui
      non placerat blandit, urna justo gravida erat, sit amet semper dolor elit
      sed metus. Mauris volutpat purus vitae sodales eleifend. Cras consequat
      scelerisque elit, sed mattis tortor lacinia non.
    </p>
		    `),
	NewExample(
		"Single panel",
		`<p>Give every <code>details</code> element in an accordion the same <code>name</code> attribute to make the sections behave as a single exclusive group, where opening one section automatically closes the section that was previously open.</p>`,
		`
  <p>
    Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nulla ullamcorper
    blandit ultricies. Etiam non suscipit felis. Orci varius natoque penatibus
    et magnis dis parturient montes, nascetur ridiculus mus. Proin feugiat
    purus hendrerit sapien condimentum, eu facilisis ex eleifend. Sed lobortis
    est a urna accumsan vestibulum. Curabitur iaculis sem lacus, id molestie
    eros ultrices eu. Nullam justo nulla, sollicitudin sed mi pretium,
    consequat varius erat. Phasellus sed eros dictum, congue enim in, dictum
    ante. Sed condimentum ut elit eu vehicula.
  </p>
  <div class="accordion">
    <details name="single">
      <summary>Section 1</summary>
      <p>
        Fusce congue nec massa id eleifend. Donec id luctus ligula. In
        pellentesque diam eu magna interdum, a tristique arcu dignissim. In eu
        lacinia nisl. Phasellus eu nisi vitae enim aliquam rutrum. Suspendisse
        fringilla tortor et tincidunt vulputate. Nam a urna a purus ornare
        gravida non sit amet risus. Morbi in justo quis velit elementum
        fermentum. Etiam tristique diam nunc, quis suscipit velit suscipit non.
      </p>
    </details>
    <details name="single">
      <summary>Section 2</summary>
      <p>
        Nam at mollis ante, non rutrum enim. Nulla facilisi. Nunc maximus diam
        a dui tincidunt vulputate. Phasellus in suscipit augue. In quis
        elementum enim. Mauris at porttitor libero. Sed vel tortor sit amet
        felis porttitor dapibus et id neque.
      </p>
    </details>
    <details name="single">
      <summary>Section 3</summary>
      <p>
        Duis facilisis eros porttitor mauris sollicitudin dapibus. Mauris
        mollis turpis nec ligula pulvinar finibus. Proin molestie varius nisl.
        Ut et ex lobortis, auctor nunc tincidunt, aliquam sem. Nunc sodales
        aliquam magna. Proin rutrum sed erat accumsan elementum. Mauris posuere
        augue non orci maximus, sed sagittis elit mattis. Proin mattis, leo at
        luctus ultrices, nulla dui bibendum ipsum, sed venenatis ex ipsum a
        erat. Ut libero dui, viverra in purus eget, efficitur pellentesque
        justo.
      </p>
    </details>
    <details name="single">
      <summary>Section 4</summary>
      <p>
        Suspendisse ultricies tristique elit eget suscipit. Sed eu odio eu diam
        convallis venenatis sit amet quis elit. Aenean mi diam, gravida eu
        nulla eget, egestas ultricies velit. Pellentesque et metus fringilla,
        lobortis felis id, suscipit nunc. Nunc interdum placerat dui, quis
        varius turpis tempor in. Maecenas faucibus tellus non lacus pulvinar, a
        malesuada urna luctus. Phasellus fermentum porta dolor vel hendrerit.
        Donec facilisis maximus ipsum congue mattis. Nullam justo arcu, aliquet
        id turpis sed, maximus laoreet dui. Nam nec elit accumsan, facilisis
        lorem ac, ullamcorper arcu. Interdum et malesuada fames ac ante ipsum
        primis in faucibus. Nullam quis rutrum libero, a sollicitudin magna.
      </p>
    </details>
  </div>
  <p>
    Pellentesque at sodales nisl, in eleifend erat. Pellentesque porttitor
    lacinia nibh, a dictum orci condimentum sed. Quisque eu maximus leo.
    Vestibulum elit arcu, molestie eget posuere vel, dapibus pharetra elit.
    Aenean eget turpis et turpis ullamcorper pretium. Praesent rutrum
    sollicitudin leo quis tincidunt. Aliquam erat volutpat. Ut viverra, dui non
    placerat blandit, urna justo gravida erat, sit amet semper dolor elit sed
    metus. Mauris volutpat purus vitae sodales eleifend. Cras consequat
    scelerisque elit, sed mattis tortor lacinia non.
  </p>
	    `))
