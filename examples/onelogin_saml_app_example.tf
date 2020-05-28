resource onelogin_saml_apps saml{
  connector_id = 50534
  name =  "SAML App"
  description = "SAML"

  configuration {
    provider_arn = null
    signature_algorithm = "SHA-1"
  }
  rules {
    enabled = true
    match = "all"
    name = "changed rule"
    position = 1
    conditions {
      operator = ">"
      source = "last_login"
      value = "90"
    }
    actions {
      action = "set_amazonusername"
      expression = ".*"
      value = ["member_of"]
    }
  }
}