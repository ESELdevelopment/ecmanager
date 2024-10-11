# Moto

[Moto](https://github.com/getmoto/moto) is a library that allows you to mock out
AWS services for testing purposes. It is a great way to test our code
without incurring the cost of running AWS services.
We use Moto as standalone server to simulate AWS locally.

## Fixtures

To modify the fixtures, we use a docker-compose sidecar, which add the fixtures to
the moto container by using Boto3. The fixtures config can be
found [here](https://github.com/ESELdevelopment/ecmanager/blob/main/compose/moto/fixture.py).

### Why do we use Python as fixtures-loader?

We use python for our local fixtures for 2 reasons:

- Boto3 is a greate SDK to interact with AWS
- Moto is written in python, if we need examples we get them on their repo.
