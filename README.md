# taslib

<center>
    <em>Trevor Arthur's personal Go package/library</em>
</center>

<hr>

### Current Functions
<h4>BasicAuth API Requests</h4>
<ul>
    <li><strong>GetApiRequestUserPass</strong><br>
        Sends a GET request to the passed in API endpoint request, using Username and Password Basic Authentication. The API JSON response is returned in a slice of bytes.</li>
    <li><strong>GetApiRequestToken</strong><br>
        Sends a GET request to the passed in API endpoint request, using Token Basic Authentication. The API JSON response is returned in a slice of bytes.</li>
    <li><strong>GetApiRequestUserPassInsecureSSL</strong><br>
        GetApiRequestUserPassInsecureSSL sends a GET request to the passed in API endpoint request,
        using Username and Password Basic Authentication, and ignore insecure SSL.
        The API JSON response is returned in a slice of bytes.</li>
    <li><strong>GetApiRequestTokenInsecureSSL</strong><br>
        GetApiRequestUserPassInsecureSSL sends a GET request to the passed in API endpoint request,
        using Token Basic Authentication, and ignore insecure SSL.
        The API JSON response is returned in a slice of bytes.</li>
</ul>

<h4>Misc</h4>
<ul>
    <li><strong>WriteToCSV</strong> - Write unmarshalled JSON to a CSV file</li>
    <li><strong>IntArrayToString</strong> - Convert an array of integers into a string</li>
</ul>
