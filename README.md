# opencost-cacher

`opencost-cacher` is a utility designed to query the OpenCost API and cache daily cost reports into a PostgreSQL database or a local SQLite database. It also exposes an API endpoint `/reports` to retrieve these cached reports based on specified query parameters.

## Features

- **Daily Cost Caching**: Automatically queries the OpenCost API daily at 3 AM and caches the reports.
- **Historical Syncing**: On startup, syncs all reports up to the latest saved report. Enable full sync to re-sync all reports regardless of existing data.
- **Database Support**: Supports caching to PostgreSQL or local SQLite databases.
- **API Endpoint**: Exposes an API endpoint `/reports` to retrieve cached reports based on time range and cluster parameters.
- **Namespace Cost Reports**: Provides detailed cost reports at the namespace level, including CPU, memory, and storage costs.

## Run Locally

```bash
git clone https://github.com/yourusername/opencost-cacher.git
cd opencost-cacher
# Set env variables
go run .
```

## Configuration

opencost-cacher is configured via environment variables. Below is a list of all configurable options:
OpenCost URLs

Set environment variables in the format OPENCOST_URL_<CLUSTER_NAME> to specify the OpenCost API URLs for each cluster.

```bash
export OPENCOST_URL_CLUSTER1=https://opencost.cluster1.example.com
export OPENCOST_URL_CLUSTER2=https://opencost.cluster2.example.com
```

### Database Configuration
#### PostgreSQL

```bash 
export DATABASE_USERNAME=your_db_username
export DATABASE_PASSWORD=your_db_password
export DATABASE_HOST=your_db_host
export DATABASE_PORT=5432
export DATABASE_DB=your_db_name
export DATABASE_SSL=disable
```

You can also set more advanced SSL settings:

```bash
export DATABASE_CA_CERT_PATH=your_ca_cert_path
export DATABASE_CLIENT_CERT_PATH=your_client_cert_path
export DATABASE_CLIENT_KEY_PATH=your_client_key_path
```
#### SQLite (Local Database)

To use a local SQLite database, set the LOCALDB environment variable:

```bash
export LOCALDB=true
```

### Sync Options

Full Sync: Enable full sync to re-sync all reports from the beginning.

```bash
export FULL_SYNC=true
```


## API Reference
/reports Endpoint

Retrieve cached namespace cost reports.
Query Parameters

    window (required): The time range for the report, specified as a from and to date in RFC3339 format, separated by a comma.
        Format: window=from_date,to_date
        Example: window=2023-01-01T00:00:00Z,2023-01-31T23:59:59Z

    cluster (required): The cluster name to filter the report. Can also use "all" to get a summarized report for all clusters.

Response Format

```json
[
    {
        "name": "namespace",
        "cpuCost": 0.0,
        "memoryCost": 0.0,
        "pVCost": 0.0,
        "totalCost": 0.0,
        "totalEfficiency": 0.0,
        "team": "team-name",
        "division": "division-name",
        "labels": {
        "key": "value"
    },
    "containers": {
        "container-name": {
            "name": "container-name",
            "cpuCost": 0.0,
            "memoryCost": 0.0,
            "pVCost": 0.0,
            "totalCost": 0.0,
            "totalEfficiency": 0.0
         }
      }
    }
]
```

    The response is a JSON array of NamespaceCost objects.
    Each NamespaceCost object contains detailed cost information for a namespace.

### Examples
Retrieve a Report for a Specific Cluster

Retrieve a report for cluster1 for the time range from February 1, 2023, to February 28, 2023.

```bash

curl "http://localhost:8080/reports?window=2023-02-01T00:00:00Z,2023-02-28T23:59:59Z&cluster=cluster1"
```

## License

This project is licensed under the MIT License.