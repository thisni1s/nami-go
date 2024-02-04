# NaMi-Go
```go
import "github.com/thisni1s/nami-go"
```

Package namigo provides a basic read-only implemaentation of the DPSG NAMI API.
It also provides a lot of structs to work with the NAMI data.

## Usage

Call ```Login("username", "password")``` to login to the NAMI. Then use the other functions to get data from it. 


#### func  GetActivityById

```go
func GetActivityById(userid string, activityId string) (types.Activity, error)
```
GetActivityById returns details about a specific "Tätigkeit". This is more
detailed then the ActivityListItem objects returned by GetMemberActivities

#### func  GetEducationById

```go
func GetEducationById(userid string, educationId string) (types.Education, error)
```
GetEducationById returns details about a specific "Ausbildung" entry. This is
more detailed then the EducationListItem objects returned by GetMemberEducation

#### func  GetGroupInfo

```go
func GetGroupInfo() (types.GroupInfo, error)
```
GetGroupInfo returns info about the primary "Stammgruppierung" from the logged
in User.

#### func  GetGroupTagById

```go
func GetGroupTagById(tagid string) (types.GroupTag, error)
```
GetGroupTagById returns details about a "Tag" of a Group This is more detailed
then the GroupTagListItem objects returned by GetGroupTags

#### func  GetGroupTags

```go
func GetGroupTags() ([]types.GroupTagListItem, error)
```
GetGroupTags returns a slice containing all "Tags" available in a "Gruppierung"

#### func  GetMemberActivities

```go
func GetMemberActivities(userid string) ([]types.ActivityListItem, error)
```
GetMemberActivities returns a slice containing all current "Tätigkeiten" of a
Member

#### func  GetMemberDetails

```go
func GetMemberDetails(userid string, groupid string) (types.Member, error)
```
GetMemberDetails gets informationa bout a specific user. This information is
more detailed then the user info contained in the SearchMember objetcs returned
by Search

#### func  GetMemberEducation

```go
func GetMemberEducation(userid string) ([]types.EducationListItem, error)
```
GetMemberEducation returns a slice containing all "Ausbildung" entries of a
Member

#### func  GetTagById

```go
func GetTagById(userid string, tagId string) (types.Tag, error)
```
GetTagById returns details about a specific "Tag" entry. This is more detailed
then the TagListEntry objects returned by GetTags

#### func  GetTags

```go
func GetTags(userid string) ([]types.TagListEntry, error)
```
GetTags returns a slice containing all "Tags" of a Member

#### func  Login

```go
func Login(username string, password string) error
```
Login authenticates with the NAMI and saves the auth cookies in a Jar to be
accessed by other functions. This function should be called before working with
any other namigo functions!

#### func  Logout

```go
func Logout() (bool, error)
```
Logout logs out the user from NAMI

#### func  Search

```go
func Search(searchValues types.SearchValues) ([]types.SearchMember, error)
```
Search searches through all Member visible to a user. Filters are specified by
values set in searchValues.

#### func  UpdateMemberActivities

```go
func UpdateMemberActivities(userid string, activity types.Activity) (int, error)
```
