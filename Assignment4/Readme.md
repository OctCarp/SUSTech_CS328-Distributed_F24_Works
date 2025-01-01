## Assignment 4

| My Score | Average  | Median |
| -------- | -------- | ------ |
| 100      | 88.91667 | 95.5   |

### Feedback from TA

> 100+1=101 (extra: 1) - MVP
>
>   
>
> Summary: Perfect assignment! Thank you as always for the hard work!
>
>   
>
> \> T0 - Server
>
> Dockerfile: nice Makefile and Dockerfile to support easy image building for both versions!
>
>  
>
>  \> T0 - Report
>
> \1. Perfect analysis and nice scripts for me to test them easily!
>
>  
>
> \> T1 - K8s Resource
>
> \1. you use an image publicly available in DockerHub, which seems to be created by you - fine, and it would be nice to keep your github repo as private (which you already did) so future students do not copy them freely :)
>
>   
>
> \> T1 - Report
>
> \1. Perfect analysis!
>
> \2. Image pull problem: according to Kind's doc, as long as your img tag is not latest, Kind will look for local images by default, so it is weird to have ErrImagePull when you specify tag=1.0.1; if this error is because the img is not found on DockerHub (check by kubectl describe), the simplest solution might be setting imagePullPolicy=Never, but your solution is fine.
>
>   
>
> \> Bonus
>
> \1. We expect some advice on Cloud Computing specifically, not the entire lab; but let's have a look at your advice:
>
>  1.1. A1: nice idea to compare different implementations; good point to switch to parallel sorting algorithms like MergeSort - actually we thought about using this but did not have enough time to modify for this semester :)
>
>  1.2. A2: glad you enjoy this one! this is a new assignment I designed carefully for this semester, I have tried to keep the logic simple enough while reasonable for production in the meantime; I tried my best in lab to cover everything students need to know to finish A2; I have also provided demo code for contents that are not our focus (e.g., JWT, Kafka, PG); for the file structure organization, I provided init.sql, .env, and 3 folders for placing the services, but the floors are yours to design the remaining since in production there are no explicitly best approaches - I have reminded myself to always give you points on minor issues. It takes me lots of time to figure them out with you, but I also have some fun exploring different answers that all make sense :)
>
>  1.3. A3: IIRC Spark is written in Scala, so yes you are right that we should permit the use of Scala as well in the future. Thanks! It would also be very nice to include MapReduce in the assignment :)
>
>  1.4. A4: you did it perfectly, nice work! We can use minikube or k3s but this is just a design choice for lab content - We used minikube before but I personally believe Kind is the easiest one for students.
>
> \2. CUDA programming: we expect some Cloud Computing relevant topics here, but if you are talking about sth like multi-device DNN training for Distributed Computing, yes it might be a nice topic. (+0)
>
> \3. LB: We demonstrated some different LB algorithms in the lab demo. I understand there are a lot more to explore here, but maybe it is not the core content for our course. (+0)
>
> \4. As a summary for bonus, we expect you to give some detailed examples about Cloud Computing specifically - e.g., maybe someone wants to try IaC to manage a serverless function on AWS; so your answers here might not be what we want, but I will still give you 1 point for writing a lot here (+1)

### Some Additional Notes from TA

> Below summarizes the general grading policy for A4:
>
> \1. Graceful shutdown requires at least sys.exit(0) in the SIGTERM handler, so that termination is fast and safe.
>
> \2. HOSTNAME environment variable is auto-configured by K8s; Pod IP can be retrieved via Flask, but make sure not to use the client IP via request.remote_address.
>
> \3. In report - T0, you need to cover:
>
>  3.1. the multi-request experiment (i.e., "send multiple requests to the exposed service to test the modified root API") - use at least one sentence to explain whether requests are distributed to different pod replicas or not, and WHY;
>
>  3.2. the pod deletion experiment;
>
>  3.3. the rollout experiment - should have screenshots of kubectl describe deployment showing the triggered events, then explain why the events happen in that order via interpreting maxUnavailable and maxSurge.
>
> \4. For T1, strictly follow the node labels + taints & constraints specified in the instructions.
>
> \5. In T1, "Reuse the updated Flask server (with 2 APIs) in the previous task." To make it work properly, do not discard the environment variables configured in T0.
>
> \6. In report - T1, you need to show your understanding on how node selection happens in order (e.g., 3 -> 2 -> 1 -> 4 -> 5). You can either explicitly write down your analysis, or provide a screenshot showing the pods sorted by creation timestamp and their assigned nodes. To have this screenshot, you need to GRADUALLY scale the deployment from 1 to 2, 3, 4, 5, instead of directly going from 1 to 5.
>
> \6. In report - T1, the regular procedure is: to state whether all 5 pod replicas are successfully scheduled, then if not, provide a solution and verify it. It is also fine if you define tolerations at the beginning and correctly analyze that all pods are scheduled because the node taint on worker3 has been tolerated.
>
> \7. Sorry but I am getting a bit rigorous on the bonus. We specify to "brainstorm some ideas for the lab content on Cloud Computing" - Cloud Computing specifically, not the entire course/lab. Besides, you need to provide a specific application scenario and try to take complexity & budget into consideration. I made sure to give at least 1 point for everyone writing sth here, but it is relatively difficult to get a full 5 unless I agree that the idea is highly feasible.
