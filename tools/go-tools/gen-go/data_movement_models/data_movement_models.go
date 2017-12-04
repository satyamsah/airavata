// Autogenerated by Thrift Compiler (1.0.0-dev)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package data_movement_models

import (
	"bytes"
	"reflect"
	"database/sql/driver"
	"errors"
	"context"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"airavata_commons"

)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = context.Background
var _ = reflect.DeepEqual
var _ = bytes.Equal

var _ = airavata_commons.GoUnusedProtection__
type DMType int64
const (
  DMType_COMPUTE_RESOURCE DMType = 0
  DMType_STORAGE_RESOURCE DMType = 1
)

func (p DMType) String() string {
  switch p {
  case DMType_COMPUTE_RESOURCE: return "COMPUTE_RESOURCE"
  case DMType_STORAGE_RESOURCE: return "STORAGE_RESOURCE"
  }
  return "<UNSET>"
}

func DMTypeFromString(s string) (DMType, error) {
  switch s {
  case "COMPUTE_RESOURCE": return DMType_COMPUTE_RESOURCE, nil 
  case "STORAGE_RESOURCE": return DMType_STORAGE_RESOURCE, nil 
  }
  return DMType(0), fmt.Errorf("not a valid DMType string")
}


func DMTypePtr(v DMType) *DMType { return &v }

func (p DMType) MarshalText() ([]byte, error) {
return []byte(p.String()), nil
}

func (p *DMType) UnmarshalText(text []byte) error {
q, err := DMTypeFromString(string(text))
if (err != nil) {
return err
}
*p = q
return nil
}

func (p *DMType) Scan(value interface{}) error {
v, ok := value.(int64)
if !ok {
return errors.New("Scan value is not int64")
}
*p = DMType(v)
return nil
}

func (p * DMType) Value() (driver.Value, error) {
  if p == nil {
    return nil, nil
  }
return int64(*p), nil
}
//Enumeration of security sshKeyAuthentication and authorization mechanisms supported by Airavata. This enumeration just
// describes the supported mechanism. The corresponding security credentials are registered with Airavata Credential
// store.
//
//USERNAME_PASSWORD:
// A User Name.
//
//SSH_KEYS:
// SSH Keys
//
//FIXME: Change GSI to a more precise generic security protocol - X509
//
type SecurityProtocol int64
const (
  SecurityProtocol_USERNAME_PASSWORD SecurityProtocol = 0
  SecurityProtocol_SSH_KEYS SecurityProtocol = 1
  SecurityProtocol_GSI SecurityProtocol = 2
  SecurityProtocol_KERBEROS SecurityProtocol = 3
  SecurityProtocol_OAUTH SecurityProtocol = 4
  SecurityProtocol_LOCAL SecurityProtocol = 5
)

func (p SecurityProtocol) String() string {
  switch p {
  case SecurityProtocol_USERNAME_PASSWORD: return "USERNAME_PASSWORD"
  case SecurityProtocol_SSH_KEYS: return "SSH_KEYS"
  case SecurityProtocol_GSI: return "GSI"
  case SecurityProtocol_KERBEROS: return "KERBEROS"
  case SecurityProtocol_OAUTH: return "OAUTH"
  case SecurityProtocol_LOCAL: return "LOCAL"
  }
  return "<UNSET>"
}

func SecurityProtocolFromString(s string) (SecurityProtocol, error) {
  switch s {
  case "USERNAME_PASSWORD": return SecurityProtocol_USERNAME_PASSWORD, nil 
  case "SSH_KEYS": return SecurityProtocol_SSH_KEYS, nil 
  case "GSI": return SecurityProtocol_GSI, nil 
  case "KERBEROS": return SecurityProtocol_KERBEROS, nil 
  case "OAUTH": return SecurityProtocol_OAUTH, nil 
  case "LOCAL": return SecurityProtocol_LOCAL, nil 
  }
  return SecurityProtocol(0), fmt.Errorf("not a valid SecurityProtocol string")
}


func SecurityProtocolPtr(v SecurityProtocol) *SecurityProtocol { return &v }

func (p SecurityProtocol) MarshalText() ([]byte, error) {
return []byte(p.String()), nil
}

func (p *SecurityProtocol) UnmarshalText(text []byte) error {
q, err := SecurityProtocolFromString(string(text))
if (err != nil) {
return err
}
*p = q
return nil
}

func (p *SecurityProtocol) Scan(value interface{}) error {
v, ok := value.(int64)
if !ok {
return errors.New("Scan value is not int64")
}
*p = SecurityProtocol(v)
return nil
}

func (p * SecurityProtocol) Value() (driver.Value, error) {
  if p == nil {
    return nil, nil
  }
return int64(*p), nil
}
//Enumeration of data movement supported by Airavata
//
//SCP:
// Job manager supporting the Portal Batch System (PBS) protocol. Some examples include TORQUE, PBSPro, Grid Engine.
//
//SFTP:
// The Simple Linux Utility for Resource Management is a open source workload manager.
//
//GridFTP:
// Globus File Transfer Protocol
//
//UNICORE_STORAGE_SERVICE:
// Storage Service Provided by Unicore
//
type DataMovementProtocol int64
const (
  DataMovementProtocol_LOCAL DataMovementProtocol = 0
  DataMovementProtocol_SCP DataMovementProtocol = 1
  DataMovementProtocol_SFTP DataMovementProtocol = 2
  DataMovementProtocol_GridFTP DataMovementProtocol = 3
  DataMovementProtocol_UNICORE_STORAGE_SERVICE DataMovementProtocol = 4
)

func (p DataMovementProtocol) String() string {
  switch p {
  case DataMovementProtocol_LOCAL: return "LOCAL"
  case DataMovementProtocol_SCP: return "SCP"
  case DataMovementProtocol_SFTP: return "SFTP"
  case DataMovementProtocol_GridFTP: return "GridFTP"
  case DataMovementProtocol_UNICORE_STORAGE_SERVICE: return "UNICORE_STORAGE_SERVICE"
  }
  return "<UNSET>"
}

func DataMovementProtocolFromString(s string) (DataMovementProtocol, error) {
  switch s {
  case "LOCAL": return DataMovementProtocol_LOCAL, nil 
  case "SCP": return DataMovementProtocol_SCP, nil 
  case "SFTP": return DataMovementProtocol_SFTP, nil 
  case "GridFTP": return DataMovementProtocol_GridFTP, nil 
  case "UNICORE_STORAGE_SERVICE": return DataMovementProtocol_UNICORE_STORAGE_SERVICE, nil 
  }
  return DataMovementProtocol(0), fmt.Errorf("not a valid DataMovementProtocol string")
}


func DataMovementProtocolPtr(v DataMovementProtocol) *DataMovementProtocol { return &v }

func (p DataMovementProtocol) MarshalText() ([]byte, error) {
return []byte(p.String()), nil
}

func (p *DataMovementProtocol) UnmarshalText(text []byte) error {
q, err := DataMovementProtocolFromString(string(text))
if (err != nil) {
return err
}
*p = q
return nil
}

func (p *DataMovementProtocol) Scan(value interface{}) error {
v, ok := value.(int64)
if !ok {
return errors.New("Scan value is not int64")
}
*p = DataMovementProtocol(v)
return nil
}

func (p * DataMovementProtocol) Value() (driver.Value, error) {
  if p == nil {
    return nil, nil
  }
return int64(*p), nil
}
// Data Movement through Secured Copy
// 
// alternativeSCPHostName:
//  If the login to scp is different than the hostname itself, specify it here
// 
// sshPort:
//  If a non-default port needs to used, specify it.
// 
// Attributes:
//  - DataMovementInterfaceId
//  - SecurityProtocol
//  - AlternativeSCPHostName
//  - SshPort
type SCPDataMovement struct {
  DataMovementInterfaceId string `thrift:"dataMovementInterfaceId,1,required" db:"dataMovementInterfaceId" json:"dataMovementInterfaceId"`
  SecurityProtocol SecurityProtocol `thrift:"securityProtocol,2,required" db:"securityProtocol" json:"securityProtocol"`
  AlternativeSCPHostName *string `thrift:"alternativeSCPHostName,3" db:"alternativeSCPHostName" json:"alternativeSCPHostName,omitempty"`
  SshPort int32 `thrift:"sshPort,4" db:"sshPort" json:"sshPort,omitempty"`
}

func NewSCPDataMovement() *SCPDataMovement {
  return &SCPDataMovement{
DataMovementInterfaceId: "DO_NOT_SET_AT_CLIENTS",

SshPort: 22,
}
}


func (p *SCPDataMovement) GetDataMovementInterfaceId() string {
  return p.DataMovementInterfaceId
}

func (p *SCPDataMovement) GetSecurityProtocol() SecurityProtocol {
  return p.SecurityProtocol
}
var SCPDataMovement_AlternativeSCPHostName_DEFAULT string
func (p *SCPDataMovement) GetAlternativeSCPHostName() string {
  if !p.IsSetAlternativeSCPHostName() {
    return SCPDataMovement_AlternativeSCPHostName_DEFAULT
  }
return *p.AlternativeSCPHostName
}
var SCPDataMovement_SshPort_DEFAULT int32 = 22

func (p *SCPDataMovement) GetSshPort() int32 {
  return p.SshPort
}
func (p *SCPDataMovement) IsSetAlternativeSCPHostName() bool {
  return p.AlternativeSCPHostName != nil
}

func (p *SCPDataMovement) IsSetSshPort() bool {
  return p.SshPort != SCPDataMovement_SshPort_DEFAULT
}

func (p *SCPDataMovement) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }

  var issetDataMovementInterfaceId bool = false;
  var issetSecurityProtocol bool = false;

  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if fieldTypeId == thrift.STRING {
        if err := p.ReadField1(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
      issetDataMovementInterfaceId = true
    case 2:
      if fieldTypeId == thrift.I32 {
        if err := p.ReadField2(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
      issetSecurityProtocol = true
    case 3:
      if fieldTypeId == thrift.STRING {
        if err := p.ReadField3(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    case 4:
      if fieldTypeId == thrift.I32 {
        if err := p.ReadField4(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  if !issetDataMovementInterfaceId{
    return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field DataMovementInterfaceId is not set"));
  }
  if !issetSecurityProtocol{
    return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field SecurityProtocol is not set"));
  }
  return nil
}

func (p *SCPDataMovement)  ReadField1(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  p.DataMovementInterfaceId = v
}
  return nil
}

func (p *SCPDataMovement)  ReadField2(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI32(); err != nil {
  return thrift.PrependError("error reading field 2: ", err)
} else {
  temp := SecurityProtocol(v)
  p.SecurityProtocol = temp
}
  return nil
}

func (p *SCPDataMovement)  ReadField3(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 3: ", err)
} else {
  p.AlternativeSCPHostName = &v
}
  return nil
}

func (p *SCPDataMovement)  ReadField4(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI32(); err != nil {
  return thrift.PrependError("error reading field 4: ", err)
} else {
  p.SshPort = v
}
  return nil
}

func (p *SCPDataMovement) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("SCPDataMovement"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(oprot); err != nil { return err }
    if err := p.writeField2(oprot); err != nil { return err }
    if err := p.writeField3(oprot); err != nil { return err }
    if err := p.writeField4(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *SCPDataMovement) writeField1(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("dataMovementInterfaceId", thrift.STRING, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:dataMovementInterfaceId: ", p), err) }
  if err := oprot.WriteString(string(p.DataMovementInterfaceId)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.dataMovementInterfaceId (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:dataMovementInterfaceId: ", p), err) }
  return err
}

func (p *SCPDataMovement) writeField2(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("securityProtocol", thrift.I32, 2); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:securityProtocol: ", p), err) }
  if err := oprot.WriteI32(int32(p.SecurityProtocol)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.securityProtocol (2) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 2:securityProtocol: ", p), err) }
  return err
}

func (p *SCPDataMovement) writeField3(oprot thrift.TProtocol) (err error) {
  if p.IsSetAlternativeSCPHostName() {
    if err := oprot.WriteFieldBegin("alternativeSCPHostName", thrift.STRING, 3); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:alternativeSCPHostName: ", p), err) }
    if err := oprot.WriteString(string(*p.AlternativeSCPHostName)); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T.alternativeSCPHostName (3) field write error: ", p), err) }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 3:alternativeSCPHostName: ", p), err) }
  }
  return err
}

func (p *SCPDataMovement) writeField4(oprot thrift.TProtocol) (err error) {
  if p.IsSetSshPort() {
    if err := oprot.WriteFieldBegin("sshPort", thrift.I32, 4); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 4:sshPort: ", p), err) }
    if err := oprot.WriteI32(int32(p.SshPort)); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T.sshPort (4) field write error: ", p), err) }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 4:sshPort: ", p), err) }
  }
  return err
}

func (p *SCPDataMovement) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("SCPDataMovement(%+v)", *p)
}

// Data Movement through GridFTP
// 
// alternativeSCPHostName:
//  If the login to scp is different than the hostname itself, specify it here
// 
// sshPort:
//  If a non-default port needs to used, specify it.
// 
// Attributes:
//  - DataMovementInterfaceId
//  - SecurityProtocol
//  - GridFTPEndPoints
type GridFTPDataMovement struct {
  DataMovementInterfaceId string `thrift:"dataMovementInterfaceId,1,required" db:"dataMovementInterfaceId" json:"dataMovementInterfaceId"`
  SecurityProtocol SecurityProtocol `thrift:"securityProtocol,2,required" db:"securityProtocol" json:"securityProtocol"`
  GridFTPEndPoints []string `thrift:"gridFTPEndPoints,3,required" db:"gridFTPEndPoints" json:"gridFTPEndPoints"`
}

func NewGridFTPDataMovement() *GridFTPDataMovement {
  return &GridFTPDataMovement{
DataMovementInterfaceId: "DO_NOT_SET_AT_CLIENTS",
}
}


func (p *GridFTPDataMovement) GetDataMovementInterfaceId() string {
  return p.DataMovementInterfaceId
}

func (p *GridFTPDataMovement) GetSecurityProtocol() SecurityProtocol {
  return p.SecurityProtocol
}

func (p *GridFTPDataMovement) GetGridFTPEndPoints() []string {
  return p.GridFTPEndPoints
}
func (p *GridFTPDataMovement) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }

  var issetDataMovementInterfaceId bool = false;
  var issetSecurityProtocol bool = false;
  var issetGridFTPEndPoints bool = false;

  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if fieldTypeId == thrift.STRING {
        if err := p.ReadField1(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
      issetDataMovementInterfaceId = true
    case 2:
      if fieldTypeId == thrift.I32 {
        if err := p.ReadField2(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
      issetSecurityProtocol = true
    case 3:
      if fieldTypeId == thrift.LIST {
        if err := p.ReadField3(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
      issetGridFTPEndPoints = true
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  if !issetDataMovementInterfaceId{
    return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field DataMovementInterfaceId is not set"));
  }
  if !issetSecurityProtocol{
    return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field SecurityProtocol is not set"));
  }
  if !issetGridFTPEndPoints{
    return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field GridFTPEndPoints is not set"));
  }
  return nil
}

func (p *GridFTPDataMovement)  ReadField1(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  p.DataMovementInterfaceId = v
}
  return nil
}

func (p *GridFTPDataMovement)  ReadField2(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI32(); err != nil {
  return thrift.PrependError("error reading field 2: ", err)
} else {
  temp := SecurityProtocol(v)
  p.SecurityProtocol = temp
}
  return nil
}

func (p *GridFTPDataMovement)  ReadField3(iprot thrift.TProtocol) error {
  _, size, err := iprot.ReadListBegin()
  if err != nil {
    return thrift.PrependError("error reading list begin: ", err)
  }
  tSlice := make([]string, 0, size)
  p.GridFTPEndPoints =  tSlice
  for i := 0; i < size; i ++ {
var _elem0 string
    if v, err := iprot.ReadString(); err != nil {
    return thrift.PrependError("error reading field 0: ", err)
} else {
    _elem0 = v
}
    p.GridFTPEndPoints = append(p.GridFTPEndPoints, _elem0)
  }
  if err := iprot.ReadListEnd(); err != nil {
    return thrift.PrependError("error reading list end: ", err)
  }
  return nil
}

func (p *GridFTPDataMovement) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("GridFTPDataMovement"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(oprot); err != nil { return err }
    if err := p.writeField2(oprot); err != nil { return err }
    if err := p.writeField3(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *GridFTPDataMovement) writeField1(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("dataMovementInterfaceId", thrift.STRING, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:dataMovementInterfaceId: ", p), err) }
  if err := oprot.WriteString(string(p.DataMovementInterfaceId)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.dataMovementInterfaceId (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:dataMovementInterfaceId: ", p), err) }
  return err
}

func (p *GridFTPDataMovement) writeField2(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("securityProtocol", thrift.I32, 2); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:securityProtocol: ", p), err) }
  if err := oprot.WriteI32(int32(p.SecurityProtocol)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.securityProtocol (2) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 2:securityProtocol: ", p), err) }
  return err
}

func (p *GridFTPDataMovement) writeField3(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("gridFTPEndPoints", thrift.LIST, 3); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:gridFTPEndPoints: ", p), err) }
  if err := oprot.WriteListBegin(thrift.STRING, len(p.GridFTPEndPoints)); err != nil {
    return thrift.PrependError("error writing list begin: ", err)
  }
  for _, v := range p.GridFTPEndPoints {
    if err := oprot.WriteString(string(v)); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err) }
  }
  if err := oprot.WriteListEnd(); err != nil {
    return thrift.PrependError("error writing list end: ", err)
  }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 3:gridFTPEndPoints: ", p), err) }
  return err
}

func (p *GridFTPDataMovement) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("GridFTPDataMovement(%+v)", *p)
}

// Data Movement through UnicoreStorage
// 
// unicoreEndPointURL:
//  unicoreGateway End Point. The provider will query this service to fetch required service end points.
// 
// Attributes:
//  - DataMovementInterfaceId
//  - SecurityProtocol
//  - UnicoreEndPointURL
type UnicoreDataMovement struct {
  DataMovementInterfaceId string `thrift:"dataMovementInterfaceId,1,required" db:"dataMovementInterfaceId" json:"dataMovementInterfaceId"`
  SecurityProtocol SecurityProtocol `thrift:"securityProtocol,2,required" db:"securityProtocol" json:"securityProtocol"`
  UnicoreEndPointURL string `thrift:"unicoreEndPointURL,3,required" db:"unicoreEndPointURL" json:"unicoreEndPointURL"`
}

func NewUnicoreDataMovement() *UnicoreDataMovement {
  return &UnicoreDataMovement{
DataMovementInterfaceId: "DO_NOT_SET_AT_CLIENTS",
}
}


func (p *UnicoreDataMovement) GetDataMovementInterfaceId() string {
  return p.DataMovementInterfaceId
}

func (p *UnicoreDataMovement) GetSecurityProtocol() SecurityProtocol {
  return p.SecurityProtocol
}

func (p *UnicoreDataMovement) GetUnicoreEndPointURL() string {
  return p.UnicoreEndPointURL
}
func (p *UnicoreDataMovement) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }

  var issetDataMovementInterfaceId bool = false;
  var issetSecurityProtocol bool = false;
  var issetUnicoreEndPointURL bool = false;

  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if fieldTypeId == thrift.STRING {
        if err := p.ReadField1(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
      issetDataMovementInterfaceId = true
    case 2:
      if fieldTypeId == thrift.I32 {
        if err := p.ReadField2(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
      issetSecurityProtocol = true
    case 3:
      if fieldTypeId == thrift.STRING {
        if err := p.ReadField3(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
      issetUnicoreEndPointURL = true
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  if !issetDataMovementInterfaceId{
    return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field DataMovementInterfaceId is not set"));
  }
  if !issetSecurityProtocol{
    return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field SecurityProtocol is not set"));
  }
  if !issetUnicoreEndPointURL{
    return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field UnicoreEndPointURL is not set"));
  }
  return nil
}

func (p *UnicoreDataMovement)  ReadField1(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  p.DataMovementInterfaceId = v
}
  return nil
}

func (p *UnicoreDataMovement)  ReadField2(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI32(); err != nil {
  return thrift.PrependError("error reading field 2: ", err)
} else {
  temp := SecurityProtocol(v)
  p.SecurityProtocol = temp
}
  return nil
}

func (p *UnicoreDataMovement)  ReadField3(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 3: ", err)
} else {
  p.UnicoreEndPointURL = v
}
  return nil
}

func (p *UnicoreDataMovement) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("UnicoreDataMovement"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(oprot); err != nil { return err }
    if err := p.writeField2(oprot); err != nil { return err }
    if err := p.writeField3(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *UnicoreDataMovement) writeField1(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("dataMovementInterfaceId", thrift.STRING, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:dataMovementInterfaceId: ", p), err) }
  if err := oprot.WriteString(string(p.DataMovementInterfaceId)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.dataMovementInterfaceId (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:dataMovementInterfaceId: ", p), err) }
  return err
}

func (p *UnicoreDataMovement) writeField2(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("securityProtocol", thrift.I32, 2); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:securityProtocol: ", p), err) }
  if err := oprot.WriteI32(int32(p.SecurityProtocol)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.securityProtocol (2) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 2:securityProtocol: ", p), err) }
  return err
}

func (p *UnicoreDataMovement) writeField3(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("unicoreEndPointURL", thrift.STRING, 3); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:unicoreEndPointURL: ", p), err) }
  if err := oprot.WriteString(string(p.UnicoreEndPointURL)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.unicoreEndPointURL (3) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 3:unicoreEndPointURL: ", p), err) }
  return err
}

func (p *UnicoreDataMovement) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("UnicoreDataMovement(%+v)", *p)
}

// LOCAL
// 
// alternativeSCPHostName:
//  If the login to scp is different than the hostname itself, specify it here
// 
// sshPort:
//  If a non-defualt port needs to used, specify it.
// 
// Attributes:
//  - DataMovementInterfaceId
type LOCALDataMovement struct {
  DataMovementInterfaceId string `thrift:"dataMovementInterfaceId,1,required" db:"dataMovementInterfaceId" json:"dataMovementInterfaceId"`
}

func NewLOCALDataMovement() *LOCALDataMovement {
  return &LOCALDataMovement{
DataMovementInterfaceId: "DO_NOT_SET_AT_CLIENTS",
}
}


func (p *LOCALDataMovement) GetDataMovementInterfaceId() string {
  return p.DataMovementInterfaceId
}
func (p *LOCALDataMovement) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }

  var issetDataMovementInterfaceId bool = false;

  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if fieldTypeId == thrift.STRING {
        if err := p.ReadField1(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
      issetDataMovementInterfaceId = true
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  if !issetDataMovementInterfaceId{
    return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field DataMovementInterfaceId is not set"));
  }
  return nil
}

func (p *LOCALDataMovement)  ReadField1(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  p.DataMovementInterfaceId = v
}
  return nil
}

func (p *LOCALDataMovement) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("LOCALDataMovement"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *LOCALDataMovement) writeField1(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("dataMovementInterfaceId", thrift.STRING, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:dataMovementInterfaceId: ", p), err) }
  if err := oprot.WriteString(string(p.DataMovementInterfaceId)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.dataMovementInterfaceId (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:dataMovementInterfaceId: ", p), err) }
  return err
}

func (p *LOCALDataMovement) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("LOCALDataMovement(%+v)", *p)
}

// Data Movement Interfaces
// 
// dataMovementInterfaceId: The Data Movement Interface has to be previously registered and referenced here.
// 
// priorityOrder:
//  For resources with multiple interfaces, the priority order should be selected.
//   Lower the numerical number, higher the priority
// 
// 
// Attributes:
//  - DataMovementInterfaceId
//  - DataMovementProtocol
//  - PriorityOrder
type DataMovementInterface struct {
  DataMovementInterfaceId string `thrift:"dataMovementInterfaceId,1,required" db:"dataMovementInterfaceId" json:"dataMovementInterfaceId"`
  DataMovementProtocol DataMovementProtocol `thrift:"dataMovementProtocol,2,required" db:"dataMovementProtocol" json:"dataMovementProtocol"`
  PriorityOrder int32 `thrift:"priorityOrder,3,required" db:"priorityOrder" json:"priorityOrder"`
}

func NewDataMovementInterface() *DataMovementInterface {
  return &DataMovementInterface{}
}


func (p *DataMovementInterface) GetDataMovementInterfaceId() string {
  return p.DataMovementInterfaceId
}

func (p *DataMovementInterface) GetDataMovementProtocol() DataMovementProtocol {
  return p.DataMovementProtocol
}

func (p *DataMovementInterface) GetPriorityOrder() int32 {
  return p.PriorityOrder
}
func (p *DataMovementInterface) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }

  var issetDataMovementInterfaceId bool = false;
  var issetDataMovementProtocol bool = false;
  var issetPriorityOrder bool = false;

  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if fieldTypeId == thrift.STRING {
        if err := p.ReadField1(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
      issetDataMovementInterfaceId = true
    case 2:
      if fieldTypeId == thrift.I32 {
        if err := p.ReadField2(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
      issetDataMovementProtocol = true
    case 3:
      if fieldTypeId == thrift.I32 {
        if err := p.ReadField3(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
      issetPriorityOrder = true
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  if !issetDataMovementInterfaceId{
    return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field DataMovementInterfaceId is not set"));
  }
  if !issetDataMovementProtocol{
    return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field DataMovementProtocol is not set"));
  }
  if !issetPriorityOrder{
    return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field PriorityOrder is not set"));
  }
  return nil
}

func (p *DataMovementInterface)  ReadField1(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  p.DataMovementInterfaceId = v
}
  return nil
}

func (p *DataMovementInterface)  ReadField2(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI32(); err != nil {
  return thrift.PrependError("error reading field 2: ", err)
} else {
  temp := DataMovementProtocol(v)
  p.DataMovementProtocol = temp
}
  return nil
}

func (p *DataMovementInterface)  ReadField3(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI32(); err != nil {
  return thrift.PrependError("error reading field 3: ", err)
} else {
  p.PriorityOrder = v
}
  return nil
}

func (p *DataMovementInterface) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("DataMovementInterface"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(oprot); err != nil { return err }
    if err := p.writeField2(oprot); err != nil { return err }
    if err := p.writeField3(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *DataMovementInterface) writeField1(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("dataMovementInterfaceId", thrift.STRING, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:dataMovementInterfaceId: ", p), err) }
  if err := oprot.WriteString(string(p.DataMovementInterfaceId)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.dataMovementInterfaceId (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:dataMovementInterfaceId: ", p), err) }
  return err
}

func (p *DataMovementInterface) writeField2(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("dataMovementProtocol", thrift.I32, 2); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:dataMovementProtocol: ", p), err) }
  if err := oprot.WriteI32(int32(p.DataMovementProtocol)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.dataMovementProtocol (2) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 2:dataMovementProtocol: ", p), err) }
  return err
}

func (p *DataMovementInterface) writeField3(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("priorityOrder", thrift.I32, 3); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:priorityOrder: ", p), err) }
  if err := oprot.WriteI32(int32(p.PriorityOrder)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.priorityOrder (3) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 3:priorityOrder: ", p), err) }
  return err
}

func (p *DataMovementInterface) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("DataMovementInterface(%+v)", *p)
}

