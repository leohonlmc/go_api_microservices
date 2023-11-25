# Go API Microservices

<p>
  <img src="https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white" style="height: 25px"/>
  <img src="https://img.shields.io/badge/MongoDB-4EA94B?style=for-the-badge&logo=mongodb&logoColor=white" style="height: 25px"/>
</p>

<table>
  <tr>
    <td style="width: 60%;">
      <h2>Project Overview</h2>
      <p>
        <code>go_api_microservices</code> is a Go-based microservices project, using Gin for routing and MongoDB for data storage. This API provides endpoints to retrieve data from <code>inventories</code> and <code>news</code> collections.
      </p>
      <h3>Features</h3>
      <ul>
        <li>RESTful API endpoints to interact with MongoDB.</li>
        <li>Utilizes Gin web framework for routing.</li>
        <li>MongoDB integration using Go driver.</li>
        <li>Clean and structured codebase.</li>
      </ul>
      
      <h2>Getting Started</h2>
      <h3>Prerequisites</h3>
      <ul>
        <li>Go (1.x or later)</li>
        <li>MongoDB</li>
        <li>Git</li>
      </ul>

      <h3>Installation</h3>
      <ol>
        <li><strong>Clone the Repository</strong></li>
        <li><strong>Set Up MongoDB</strong>
          <ul>
            <li>Ensure MongoDB is running.</li>
            <li>Use the <code>inventories</code> and <code>news</code> collections.</li>
          </ul>
        </li>
        <li><strong>Configure Environment Variables</strong>
          <p>Create a <code>.env</code> file in the root with your MongoDB URI:</p>
        </li>
        <li><strong>Run the Application</strong></li>
      </ol>

      <h2>Usage</h2>
      <h3>API Endpoints</h3>
      <ul>
        <li><strong>Get Inventories</strong>
          <p><code>GET /inventories</code> - Fetches documents from the <code>inventories</code> collection.</p>
        </li>
        <li><strong>Get News</strong>
          <p><code>GET /news</code> - Retrieves documents from the <code>news</code> collection.</p>
        </li>
      </ul>

      <h3>Example Request</h3>
      <p>Using <code>curl</code> to fetch inventories:</p>

      <h2>Contributing</h2>
      <p>Contributions are welcome. To contribute:</p>
      <ol>
        <li>Fork the repository.</li>
        <li>Create a new branch (<code>git checkout -b my-new-feature</code>).</li>
        <li>Commit your changes (<code>git commit -am 'Add some feature'</code>).</li>
        <li>Push to the branch (<code>git push origin my-new-feature</code>).</li>
        <li>Create a new Pull Request.</li>
      </ol>
    </td>
    <td>
      <img src="./go.png" alt="Go Bear Logo" width="400" height="200">
    </td>

  </tr>
</table>
