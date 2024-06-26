package model

type Action string

const (
   Create Action = "Create"
   Update Action = "Update"
   Delete Action = "Delete"
   Read   Action = "Read"
   List   Action = "List"
)

type Model interface {
   // GetResourceModel Return the generated resource.Model
   GetResourceModel() interface{}
   // GetResourceModels Return the generated resource.Model(s) as an array for List
   GetResourceModels() []interface{}
   // GetGraphQLFragment get the GraphQL fragment from resource.Model
   GetGraphQLFragment() *string
   // SetIdentifier set the identifier in the generated resource.Model
   SetIdentifier(g *string)
   // GetIdentifier get the identifier in the generated resource.Model
   GetIdentifier() *string
   // GetIdentifierKey the response key that has the guid/id/pk
   GetIdentifierKey(a Action) string
   // GetCreateMutation get the GraphQL mutation for Create
   GetCreateMutation() string
   // GetDeleteMutation get the GraphQL mutation for Delete
   GetDeleteMutation() string
   // GetUpdateMutation get the GraphQL mutation for Update
   GetUpdateMutation() string
   // GetReadQuery get the GraphQL query for Read
   GetReadQuery() string
   // GetListQuery get the GraphQL query for List
   GetListQuery() string
   // GetListQueryNextCursor get the GraphQL query for pagination, NEXTCURSOR is the template param
   GetListQueryNextCursor() string
   // GetVariables return moustache substitution variables from resource.Model
   GetVariables() map[string]string
   // AppendToResourceModels adds a resource.Model to resource.Model.ResourceModels
   AppendToResourceModels(m Model)
   // NewModelFromGuid creates a new Model with the passed guid TODO change Guid to Identifier
   NewModelFromGuid(g interface{}) Model
   // GetErrorKey returns the key of the error 'type' field
   GetErrorKey() string
   HasTags() bool
   GetTags() map[string]string
   GetTagIdentifier() *string
}
