<!--
order: 0
title: notifications Overview
parent:
  title: "notifications"
-->
[◀ modules](/x/README.md)

# `notifications`

## Contents
1. [Concept](#concept)
2. [Client](#client)
    + [Query](#query)
    + [Transactions](#transactions)
3. [Transaction Messages](#transaction-messages)
    + [SetCounter](#setcounter)
    + [CreateNotifications](#createnotifications)
    + [DeleteNotifications](#deletenotifications)
    + [BlockSenders](#blocksenders)
4. [Query Requests](#query-requests)
    + [GetNotiCounter](#getnoticounter)
    + [AllNoticationsByAddress](#allnotificationsbyaddress)


## Concept
The `notifications` module is responsible for sending users notifications. It complements the filetree module to enable file sharing.  

## Client

Below are CLI query and transaction commands to interact with canined.
### Query
The `query` commands allow users to query `notifications` state.
```sh
canined q notifications --help
```

### Transactions
The `tx` commands allow users to interact with the `notifications` module.
```sh
canined tx notifications --help
```

## Transaction Messages

Below is a full description of valid transaction messages that can be broadcasted to affect state change. These descriptions aim to be "implementation agnostic", i.e., they make sense to both the CLI/Golang and TS implementations. 

### SetCounter

Set a NotiCounter object on chain that keeps track of how many notifications a user has. Alice will NOT be able to create a notification for Bob unless Bob has set a notiCounter. 

|Name|Type|Description|                                                                                       
|--|--|--|
|creator  | String  | The creator and broadcaster of this message. Pass in Bech32 address 

#### Response

|Name|Type|Description|                                                                                       
|--|--|--|
|NotiCounter  | String  | The actual count of how many notifications a user has. This will return 0.  

### CreateNotifications

Create a notification for a recipient 

|Name|Type|Description|                                                                                       
|--|--|--|
|creator  | String  | The creator and broadcaster of this message. Pass in Bech32 address<br /> 
|notification  | String  | The notification itself. <br /> Please note that the notification keeper does not handle encryption of this string.<br /> 
|address  | String  | The recipient address. Pass in Bech32 address<br />

#### Response

|Name|Type|Description|                                                                                       
|--|--|--|
|NotiCounter  | String  | The actual count of how many notifications a user has.

### DeleteNotifications

This message will delete the latest notification from the caller's notifications list.

|Name|Type|Description|                                                                                       
|--|--|--|
|creator  | String  | The creator and broadcaster of this message. Pass in Bech32 address<br /> 

#### Response

|Name|Type|Description|                                                                                       
|--|--|--|
|NotiCounter  | String  | The actual count of how many notifications a user has.

### BlockSenders

|Name|Type|Description|                                                                                       
|--|--|--|
|creator  | String  | The creator and broadcaster of this message. Pass in Bech32 address<br /> 
|senderIds  | String  | a stringified array of Bech32 addresses the user wants to prevent from sending them notifications. E.g., ["addressA","addressB","addressC"] <br />

#### Response

Coming soon. For now, an empty response with no errors is a good response. 

## Query Requests

Below is a full description of valid query requests that can be made to retrieve state information. These descriptions aim to be "implementation agnostic", i.e., they make sense to both the CLI/Golang and TS implementations.

### GetNotiCounter

Retrieve a user's NotiCounter

|Name|Type|Description|                                                                                       
|--|--|--|
|address  | String  | user's Bech32 address<br />

#### Response

types.NotiCounter

### AllNotificationsByAddress

Retrieve all notifications for a given address

|Name|Type|Description|                                                                                       
|--|--|--|
|address  | String  | user's Bech32 address<br />



#### Response

[]types.Notifications


