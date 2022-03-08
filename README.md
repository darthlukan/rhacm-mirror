# rhacm-mirror

Author: Brian Tomlinson <btomlins@redhat.com>

## Description

An OpenShift/Kubernetes operator which orchestrates automatic failover of one RHACM hub to another in a multi-hub environment.

### Use case

RHACM hub clusters are not impervious to outage:
1. Data centers may lose connectivity
2. Hardware may fail
3. Cloud providers may come under attack
4. Operations teams may not realize they pushed a bad commit that shuts down a cluster
5. etc

Organizations running a single hub cluster may find themselves in the midst of an outage and scrambling to recover the hub in order to maintain CI/CD throughput
to their managed clusters for RHACM managed applications and integrations, or to update managed cluster configuration state for a critical change window. 

Even mature IT organizations operating multiple hubs run into the problem of having to manually trigger automation in order to delegate reconciliation of 
desired state from git to managed clusters. This takes valuable time and due to the risk of human error, more time than is ever planned, especially if your
operations team relies on manual runbooks rather than automation.

The RHACM Mirror operator seeks to solve this problem by taking advantage of the backup features provided in RHACM 2.4+ and automatically orchestrate leadership of
reconciliation to the next available hub in the "hub pool", making manual intervention unnecessary.

### Solution

The solution to this problem is simple: RHACM Mirror.

The RHACM Mirror operator is deployed to all hub clusters. A single CR is applied as configuration: `kind: HubPool` which lists the API URLs for each of those
hubs (`spec.pool: [{}]`) and an initial "leader" hub (`spec.initial_leader`).

RHACM Mirror sends a `GET` request periodically, every 2 minutes (tuneable, `spec.ping_period`), to each Hub API URL in the pool. If the `initial_leader` does not
respond 3 times (tuneable, `spec.ping_fails_accepteptable`) during `ping_period` then RHACM Mirror will select the next hub in the pool as the new leader, restore
the latest backups from the configured backup destination (`spec.backup_location: {}`) by pulling them in and "applying" (equivalent of `oc apply -f ...`) them to
the new leader, and send an alert to the designated communication channel (`spec.failover_alert_email`, more channels to be developed later) notifying of the
change.

*EXAMPLE:*
```
apiVersion: rhacm-mirror.openshift.io/v1alpha1
kind: HubPool
metadata:
  name: my-hubpool
  namespace: rhacm-mirror
  labels:
    app: rhacm-mirror
spec:
  pool:
    - hub0: api.hub0.example.com
    - hub1: api.hub1.example.com
    - hub2: api.hub2.example.com

  initial_leader: hub0

  ping_period: 2m

  ping_fails_acceptable: 3

  backup_location:
    url: hub-backup.example.com
    protocol: s3
    credentials:
      secret: secret-containing-creds-to-s3

  failover_alert_email: ops-admin@example.com
```

With the above config, `hub0` is the initial leader. If `hub0` fails to respond 3 times in 6 minutes (3fails * 2minutes == 6 minutes) to `hub1` and `hub2`, after 
the third failure RHACM Mirror running on `hub1` will select `hub1` as the new leader, `hub2` will "agree", and RHACM Mirror running on `hub1` will begin
restoring the backup at `hub-backup.example.com` to itself in order to reconnect to the managed clusters and the git repos that back them.


## Deploy RHACM Mirror

TBD (RHACM [Application + Subscription + PlacementRule + Channel] kustomize templates for deployment via RHACM to RHACM)


## Requirements

TBD


## LICENSE

MIT
