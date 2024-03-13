# Elastic APM

Qualitytrace fetches traces from [Elasticsearch's default port](https://discuss.elastic.co/t/what-are-ports-9200-and-9300-used-for/238578) `9200`.

:::tip
Examples of configuring Qualitytrace can be found in the [`examples` folder of the Qualitytrace GitHub repo](https://github.com/intelops/quality-trace/tree/main/examples).
:::

## Configure Qualitytrace to Use Elastic APM as a Trace Data Store

Configure Qualitytrace to fetch trace data from Elasticsearch.

Qualitytrace uses Elasticsearch's **default port** `9200` to fetch trace data.

You need to know which **Index name**, **Address**, and **credentials** you are using.

The defaults can be:

- **Index name**: `traces-apm-default`
- **Address**: `https://es01:9200`
- **Username**: `elastic`
- **Password**: `changeme`

To configure Elastic APM you will need to download the CA certificate from the Docker image and upload it to the config under "Upload CA file".

- The command to download the `ca.crt` file is:
  `docker cp quality-trace-elasticapm-with-elastic-agent-es01-1:/usr/share/elasticsearch/config/certs/ca/ca.crt .`
- Alternatively, you can skip CA certificate validation by setting the `Enable TLS but don't verify the certificate` option.

:::tip
Need help configuring the OpenTelemetry Collector so send trace data from your application to Elastic? Read more in [the reference page here](../opentelemetry-collector-configuration-file-reference).
:::

## Connect Qualitytrace to Elastic with the Web UI

In the Web UI, (1) open Settings and, on the (2) Configure Data Store tab, select (3) Elastic APM. If you are using Docker, as in the screenshot below, use the service name as the hostname with port `9200`. Use `http`, or `https` if TLS is enabled.

```
https://es01:9200
```

![ElasticAPM](../img/ElasticAPM-settings.png)

<!---![](https://res.cloudinary.com/djwdcmwdz/image/upload/v1674566041/Blogposts/Docs/screely-1674566018046_ci0st9.png)-->

## Connect Qualitytrace to Elastic with the CLI

Or, if you prefer using the CLI, you can use this file config.

```yaml
type: DataStore
spec:
  name: Elastic Data Store
  type: elasticapm
  default: true
  elasticapm:
    addresses:
      - https://es01:9200
    username: elastic
    password: changeme
    index: traces-apm-default
    insecureSkipVerify: true
```

Proceed to run this command in the terminal and specify the file above.

```bash
quality-trace apply datastore -f my/data-store/file/location.yaml
```

:::tip
To learn more, [read the recipe on running a sample app with Elastic APM and Qualitytrace](../../examples-tutorials/recipes/running-quality-trace-with-elasticapm.md).
:::
