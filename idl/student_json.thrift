// idl/student.thrift
namespace go student

struct Student {
    1: required string studentNo(api.body='studentNo'),
    2: required string studentName(api.body='studentName'),
}

struct QueryStudentReq {
    1: string studentNo (api.body="studentNo"); // 添加 api 注解为方便进行参数绑定
}

struct Response {
    1: i32 Code;
    2: string Msg;
    3: string Data;
}

struct SaveStudentReq {
    1: string StudentNo (api.body="studentNo");
    2: string StudentName (api.body="studentName");
}


service StudentService {
    Response AddStudent(1: SaveStudentReq request) (api.post="/add-student-info");
    Response QueryStudent(1: QueryStudentReq request) (api.get="/query");
}
