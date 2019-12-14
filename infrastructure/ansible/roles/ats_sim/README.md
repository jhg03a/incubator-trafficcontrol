<!--
    Licensed to the Apache Software Foundation (ASF) under one
    or more contributor license agreements.  See the NOTICE file
    distributed with this work for additional information
    regarding copyright ownership.  The ASF licenses this file
    to you under the Apache License, Version 2.0 (the
    "License"); you may not use this file except in compliance
    with the License.  You may obtain a copy of the License at

      http://www.apache.org/licenses/LICENSE-2.0

    Unless required by applicable law or agreed to in writing,
    software distributed under the License is distributed on an
    "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
    KIND, either express or implied.  See the License for the
    specific language governing permissions and limitations
    under the License.
-->
ats_sim
=========



Requirements
------------



Role Variables
--------------

Refer to the defaults/main.yml for most information.

Dependencies
------------

None

Example Playbook
----------------
```yaml
  - name: Deploy ats_sim
    import_role:
      name: ats_sim
    vars:
      install_ats_sim: true
      ats_sim_to_url: {{ to_url }}
      ats_sim_to_user: {{ to_user }}
      ats_sim_to_password: {{ to_password }}
      ats_sim_to_cdn: Kabletown2.0
      ats_sim_git_repo: "https://github.com/apache/trafficcontrol.git"
      ats_sim_git_version: master
```

License
-------

Apache 2.0

Author Information
------------------

Apache Traffic Control
