tot

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
  * Sip codes
  * guid

## Usage

#### Create config file

    ; Default api-config
    ; https://github.com/robfig/config
    [DEFAULT]
    host: api.tropo.com
    route: /v1
    protocol: https://
    base-url: %(protocol)s%(host)s%(route)s

    [hosted]
    url: %(base-url)s
    username: username
    password: abc123

##### Install Application
    go get github.com/mattn/gom
    gom install
    gom build
    mv tropo /bin/tropo && chmod +x /bin/tropo
    tropo

##### Rule the world

    NAME:
       tot - Tropo operations utility

    USAGE:
       tot [global options] command [command options] [arguments...]

    VERSION:
       0.2.0

    COMMANDS:
       guid     tropo guid 9fb9f0887171a133e4ce14025baa968e
       list     tropo list features
       lookup   tropo lookup 9fb9f0887171a133e4ce14025baa968e
       help, h  Shows a list of commands or help for one command

    GLOBAL OPTIONS:
       --config, -c 'hosted'    config to use in ~/.tropo-api.cfg
       --version, -v        print the version
       --help, -h           show help


### To do

* Handle auth errors better, currently its a giant stacktrace
* Add payment info https://api.tropo.com/v1/users/jdyer/payment/method
##### Lookup data


###### Lookup Features

    $ tot list features                                                                                                                                                                                                                             ±[●][master]
    +----------+------------------------------+
    | PROPERTY |            VALUE             |
    +----------+------------------------------+
    | s        | Outbound SIP                 |
    | b        | SIP Bang Syntax              |
    | c        | Override Caller id           |
    | w        | International Outbound SMS   |
    | i        | Internation    al Outbound Voice |
    | u        | Domestic Outbound Voice      |
    | d        | Domestic Outbound SMS        |
    | r        | SIP REFER                    |
    | x        | Disabled Account             |
    +----------+------------------------------+

###### Decode sessionGUID

     $ tot guid 9fb9f0887171a133e4ce14025baa968e
     Results: 10.6.69.185

###### Address

    $ tot lookup address +14433058696
    Results
    +---------------------+-----------------+
    |      PROPERTY       |      VALUE      |
    +---------------------+-----------------+
    | Type                | number          |
    | Prefix              | 1510            |
    | Number              | +151021012345    |
    | DisplayNumber       | +1 510-210-12345 |
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

    $ tot lookup user jdyer
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

    $ tot lookup application 5032272
    Results
    +--------------+--------------------------------------------------------------------+
    |   PROPERTY   |                               VALUE                                |
    +--------------+--------------------------------------------------------------------+
    | AppId        | 1234566                                                            |
    | UserId       | 12345                                                              |
    | App Name     | ppid832                                                            |
    | Platform     | scripting                                                          |
    | Environment  | http://api.tropo.com/v1/environments/461                  |
    | MessagingUrl | https://dl.dropboxusercontent.com/u/1234/TROPO/app.rb              |
    | VoiceUrl     | https://dl.dropboxusercontent.com/u/1234/TROPO/appMsg.rb         |
    | Partition    | production                                                         |
    +--------------+--------------------------------------------------------------------+

### Multiple API Support


Tropo can support multiple api endpoints.  You simply need to configure new sections as illustrated below:
#### Config

    [DEFAULT]
    host: api.tropo.com
    route: /v1
    protocol: https://
    base-url: %(protocol)s%(host)s%(route)s

    [hosted]
    url: %(base-url)s
    username: jdyer
    password: abc

    [testing]
    url: %(protocol)stesting.tropo.com%(route)s
    username: jdyer
    password: abc

    [development]
    url: %(protocol)development.tropo.com%(route)s
    username: jdyer
    password: abc


Once this is done you can perform the same actions but use a global flag to point to the different backend

    $ tot --config testing lookup user jdyer                                                                                                                                                                                                                                       ±[master]
    Results
    +------------------------+------------------------------+
    |        PROPERTY        |            VALUE             |
    +------------------------+------------------------------+
    | Username               | jdyer                        |
    | AccountId              | 4                            |
    | Email                  | nobody@nowhere.com           |
    | Name                   | n/a n/a                      |
    | Address                |                              |
    | JoinDate               | 2012-06-08T20:35:00.000+0000 |
    | Status                 | locked                       |
    | PasswordFailedAttempts | 6                            |
    | Feature Flags          | s,c,u                        |
    +------------------------+------------------------------+



##### Disable SSL validation

For some hosts, especially development ones, you want to use a self signed certificate.  This can be done by simply adding the following flag globally or per backend

    [development]
    url: %(protocol)splayground.tropo.com%(route)s
    username: jdyer
    password: abc
    InsecureSkipVerify: true

## Author

[John Dyer]()

=======
