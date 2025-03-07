# Zoom Meeting Integration

This project is bootstraped with NX monorepo. Use the NX plugin on VS Code, Cursor, or Trae IDE.

## How To Use

### Install

Install all the dependencies with command below

```bash
yarn
```

### Run the Api

Start the backend or apis by running the command below

```bash
ZOOM_ACCOUNT_ID={account_id} ZOOM_AUTHORIZATION={base64token} ZOOM_REQUEST_TOKEN_URL=https://zoom.us/oauth/token go run apps/b7-zoom-integration-apis/main.go
```

### Run the Web App

Start the frontend by running the command below

```bash
yarn nx run b7-zoom-integration:dev
```
