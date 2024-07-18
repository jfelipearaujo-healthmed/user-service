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
- ✅: Completed
- 🚧: In progress
- 💤: Not started


| Completed | Method | Endpoint                       | Description                  |
| --------- | ------ | ------------------------------ | ---------------------------- |
| ✅         | GET    | `/users/me`                    | Get the current user         |
| ✅         | POST   | `/users`                       | Create a user                |
| ✅         | PUT    | `/users/me`                    | Update a user                |
| 💤         | GET    | `/users/me/reviews`            | Get the current user reviews |
| 🚧         | POST   | `/users/me/reviews`            | Create a user review         |
| 💤         | GET    | `/users/me/reviews/{reviewId}` | Get a user review by id      |

# License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.