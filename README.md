# User Service

Service responsible to manage the users

# Local Development

## Requirements

- [Kubernetes](https://kubernetes.io/)
- [AWS CLI](https://aws.amazon.com/cli/)

## Manual deployment

### Attention

Before deploying the service, make sure to set the `AWS_ACCESS_KEY_ID` and `AWS_SECRET_ACCESS_KEY` environment variables.

Be aware that this process will take a few minutes (~4 minutes) to be completed.

To deploy the service manually, run the following commands in order:

```bash
make init
make check # this will execute fmt, validate and plan
make apply
```

To destroy the service, run the following command:

```bash
make destroy
```

## Automated deployment

The automated deployment is triggered by a GitHub Action.

# Endpoints

Legend:
- ✅: Development completed
- 🚧: In progress
- 💤: Not started


| Completed | Method | Endpoint                    | Description                               | User Role      |
| --------- | ------ | --------------------------- | ----------------------------------------- | -------------- |
| ✅         | POST   | `/users`                    | Create a user                             | Doctor/Patient |
| ✅         | GET    | `/users/me`                 | Get the current user                      | Doctor/Patient |
| ✅         | PUT    | `/users/me`                 | Update a user                             | Doctor/Patient |
| ✅         | GET    | `/users/doctors`            | Get doctors by Medical ID, specialty, etc | Patient        |
| ✅         | GET    | `/users/doctors/{doctorId}` | Get doctor by ID                          | Patient        |


# Diagrams

## User Handling

In this diagram, we can see all the possible interactions between the user and the service.

Attention: The user could be a doctor or a patient.

![user_handling](./docs/user_creation.svg)

# License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.