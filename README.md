tropo

====

## Synopsis

Small go library used by Tropo team to lookup data from Tropo Provisioning API.

Types of lookups:

  * address
    * pin
    * token
    * number
  * users
    * username
    * userId
  * Applications
    * applicationId
  * guid

## Usage

#### Create config file

    [api]
    url = https://api.aws.tropo.com/rest/v1
    [credentials]
    username = jdyer
    password = abc123

##### Install Application

    go install



##### Lookup data

###### Decode sessionGUID

     $ tropo guid 9fb9f0887171a133e4ce14025baa968e
     Results: 10.6.69.185

###### Address

    tropo address -n +15102101549
    Results
    +---------------------+-----------------+
    |      PROPERTY       |      VALUE      |
    +---------------------+-----------------+
    | Type                | number          |
    | Prefix              | 1510            |
    | Number              | +15102101549    |
    | DisplayNumber       | +1 510-210-1549 |
    | ServiceId           | 5150065         |
    | City                | Crockett        |
    | State               | CA              |
    | Country             | United States   |
    | ProviderName        | bw              |
    | SmsEnabled          | true            |
    | ExcludeFromBilling  | false           |
    | SmsRateLimit        | 60              |
    | ExchangeId          | 2584            |
    | ApplicationId       | 5032096         |
    | RequireVerification | false           |
    +---------------------+-----------------+

###### User

    tropo user -u jdyer
    Results
    +------------------------+--------------------------------+
    |        PROPERTY        |             VALUE              |
    +------------------------+--------------------------------+
    | Username               | jdyer                          |
    | AccountId              | 29890                          |
    | Email                  | jdyer@voxeolabs.com            |
    | Name                   | John Dyer                      |
    | Address                | 189 South Orange Avenue #1000  |
    |                        | FL                             |
    | JoinDate               | 2008-04-05T19:57:50.000+0000   |
    | Status                 | active                         |
    | Notes                  | none                           |
    | PasswordFailedAttempts | 0                              |
    +------------------------+--------------------------------+

###### Application

    $ tropo application -a 5032272
    Results
    +--------------+--------------------------------------------------------------------+
    |   PROPERTY   |                               VALUE                                |
    +--------------+--------------------------------------------------------------------+
    | AppId        | 1234566                                                            |
    | UserId       | 12345                                                              |
    | App Name     | ppid832                                                            |
    | Platform     | scripting                                                          |
    | Environment  | http://api.aws.tropo.com/rest/v1/environments/461                  |
    | MessagingUrl | https://dl.dropboxusercontent.com/u/1234/TROPO/app.rb              |
    | VoiceUrl     | https://dl.dropboxusercontent.com/u/177285/TROPO/appMsg.rb         |
    | Partition    | production                                                         |
    +--------------+--------------------------------------------------------------------+

## Author

[John Dyer]()

=======
