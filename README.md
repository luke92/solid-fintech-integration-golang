# solid-fintech-integration-golang
Integration with Solid Fintech 

# RUN
make -f MakeFile

## Configuration PORT and SOLID
- Create the file `dev.env` and save inside the folder `pkg/common/config/envs/`

## File dev.env
```
PORT=:3000
SOLID_ENV=https://test-api.solidfi.com
SOLID_API_KEY=YOUR_API_KEY
VGS_API=https://tntbevlgikb.sandbox.verygoodproxy.com 
```

## VGS
Very good security is used for set new PIN for a card

## Documentation of Solid Fintech
https://www.solidfi.com/docs/dashboard

